using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Net;
using System.Net.Sockets;
using System;
using static TcpClient;
using System.IO;
using System.Threading;

//定义回调接口给外部使用
public interface ISocketCallBack
{
    void OnConnect();  //链接完成后给外部
    void OnDisconnect();  //链接断开后给外部
    void OnReceive(byte[] data, int len);
}

public class TcpClient
{
    public enum ConnectStatus
    {
        INVALID = 0,  //无效的
        CONNECTING,  //正在链接
        CONNECT_FINISH,  //链接完成
        CONNECTED,  //链接已成功
        CONNECT_FAIL, //链接失败
        DISCONNECTED,    //已断联
    }

    private ConnectStatus m_SocketStatus;

    private Socket m_Socket = null;
    private byte[] m_ReceiveData = new byte[4096];

    private ArrayBuffer mReceiveBuffer = new ArrayBuffer();
    private Mutex mMutex = new Mutex();

    ISocketCallBack mCb = null;

    private bool m_Init = false;
    public void Init(ISocketCallBack callBack)
    {
        if (m_Init)
        {
            return;
        }
        mCb = callBack;
        m_Socket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
    }

    //建立异步链接
    public void Connect(string host, int port)
    {
        m_Socket.BeginConnect(host, port, new AsyncCallback(OnAsyncConnect), null);
        m_SocketStatus = ConnectStatus.CONNECTING;

        Debug.Log(string.Format("连接线程0{0}", Thread.CurrentThread.ManagedThreadId.ToString()));
    }

    public void OnAsyncConnect(IAsyncResult ar)
    {
        Debug.Log(string.Format("连接线程1{0}", Thread.CurrentThread.ManagedThreadId.ToString()));
        try
        {
            m_Socket.EndConnect(ar);
            if (!m_Socket.Connected)
            {
                m_SocketStatus = ConnectStatus.CONNECT_FAIL;
                return;
            }

            m_SocketStatus = ConnectStatus.CONNECT_FINISH;
            m_Socket.BeginReceive(m_ReceiveData, 0, 4096, SocketFlags.None, new AsyncCallback(OnAsyncReceive), null);
        }
        catch (Exception ex)
        {
            Debug.LogException(ex);
        }
    }

    public void OnAsyncReceive(IAsyncResult ar)
    {
        try
        {
            int length = m_Socket.EndReceive(ar);

            //lock
            mMutex.WaitOne();
            mReceiveBuffer.WriteBytes(m_ReceiveData, length);
            mMutex.ReleaseMutex();

            //继续接收数据
            m_Socket.BeginReceive(m_ReceiveData, 0, 4096, SocketFlags.None, new AsyncCallback(OnAsyncReceive), null);
        }
        catch (Exception ex)
        {
            Debug.LogException(ex);
        }
    }

    public void Send(byte[] data, int len)
    {
        try
        {
            byte[] buffer = new byte[len + 4];
            BytesHelper.Int32ConvertByte(len, buffer, 4, 0);
            Array.Copy(data, 0, buffer, 4, len);
            m_Socket.BeginSend(buffer, 0, buffer.Length, SocketFlags.None, new AsyncCallback(OnAsyncSend), null);
        }
        catch (Exception ex)
        {
            Debug.LogException(ex);
        }
    }

    public void OnAsyncSend(IAsyncResult ar)
    {
        m_Socket.EndSend(ar);
    }

    private bool IsValid()
    {
        if (m_Socket == null)
        {
            return false;
        }
        return true;
    }

    private bool IsConnect()
    {
        if (!IsValid())
        {
            return false;
        }
        if (m_Socket.Connected)
        {
            return true;
        }
        if (m_SocketStatus == ConnectStatus.CONNECTED)
        {
            return true;
        }
        return false;
    }

    public void Update()
    {
        if (m_SocketStatus == ConnectStatus.CONNECT_FINISH)
        {
            m_SocketStatus = ConnectStatus.INVALID;
            OnConnect();
        }
        if (!IsConnect())
        {
            return;
        }
        ProcessSocket();
    }

    private void OnConnect()
    {
        m_SocketStatus = ConnectStatus.CONNECTED;
        mCb.OnConnect();
    }

    private void OnClose()
    {
        m_SocketStatus = ConnectStatus.INVALID;
        mCb.OnDisconnect();
    }

    private void ProcessSocket()
    {
        try
        {
            if (!IsConnect())
            {
                return;
            }
            ProcessInput();
        }
        catch (Exception ex)
        {
            Debug.LogException(ex);
            OnClose();
        }
    }

    //处理网络中接收到的数据
    private void ProcessInput()
    {
        try
        {
            mMutex.WaitOne();
            if (mReceiveBuffer.Length <= 0)
            {
                mMutex.ReleaseMutex();
                return;
            }
            byte[] buf=new byte[mReceiveBuffer.Length];
            mReceiveBuffer.Read(buf);
            mReceiveBuffer.Seek(0);
            mMutex.ReleaseMutex();

            mCb.OnReceive(buf, buf.Length);
        }
        catch (Exception ex)
        {
            mMutex.ReleaseMutex();
            Debug.LogException(ex);
        }
    }
}


