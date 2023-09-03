using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Net;
using System.Net.Sockets;
using System;
using static TcpClient;
using System.IO;

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
        CONNECT_FAIL  //链接失败
    }

    private ConnectStatus m_Status;

    private Socket m_Socket = null;
    private SocketInputStream m_InputStream;
    private SocketOutputStream m_OutputStream;

    private bool m_Init = false;
    public void Init(ISocketCallBack callBack)
    {
        if (m_Init)
        {
            return;
        }
        m_Socket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
        m_InputStream = new SocketInputStream(this);
        m_OutputStream = new SocketOutputStream(this);
    }

    //建立同步链接
    public void Connect(string host, int port)
    {
        m_Socket.BeginConnect(host, port, new AsyncCallback(OnConnectCallBack), this);
        m_Status = ConnectStatus.CONNECTING;
    }

    static public void OnConnectCallBack(IAsyncResult ar)
    {
        TcpClient client = (TcpClient)ar.AsyncState;

        client.OnConnect();
    }

    private void OnConnect()
    {
        if (m_Socket.Connected)
        {
            m_Status = ConnectStatus.CONNECTED;
        }
        else
        {
            m_Status = ConnectStatus.CONNECT_FAIL;
        }
    }

    private bool IsValid()
    {
        if (m_Socket == null)
        {
            return false;
        }
        if (m_Socket.Available == 0)
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
            return false;
        }
        if (m_Status == ConnectStatus.CONNECTED)
        {
            return true;
        }
        return false;
    }

    public void Update()
    {
        if (!IsConnect())
        {
            return;
        }
        ProcessSocket();
    }

    ////发送数据
    //public void Send(byte[] data, int len)
    //{
    //    m_Socket.Send(data, len, SocketFlags.None);
    //}

    ////接收数据
    //public void Receive(byte[] data)
    //{
    //    int dataLen = m_Socket.Receive(data);
    //}

    private void ProcessSocket()
    {
        ProcessInput();

        ProcessOutput();
    }

    //处理网络中接收到的数据
    private void ProcessInput()
    {
        //从Socket中的Stream中取出数据
        try
        {
            if (m_InputStream == null)
            {
                return;
            }
            if (!IsConnect())
            {
                return;
            }
            m_InputStream.Fill();
        }
        catch (Exception e)
        {
            LogManager.Instance.ErrorLog("socket error" + e.ToString());
        }
    }

    //处理要发送的数据
    private void ProcessOutput()
    {
        //把数据写入Socket
        try
        {
            if (m_OutputStream == null)
            {
                return;
            }
            if (!IsConnect())
            {
                return;
            }
            m_OutputStream.Flush();
        }
        catch (Exception e)
        {
            LogManager.Instance.ErrorLog("socket error" + e.ToString());
        }
    }

    public Socket GetSocket()
    {
        return m_Socket;
    }

    public void Send(byte[] data, int len)
    {
        m_OutputStream.Write(data, len);
    }
}


