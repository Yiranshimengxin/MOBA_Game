using System;
using System.Collections;
using System.Collections.Generic;
using System.Net.Sockets;
using UnityEngine;

public class SocketOutputStream
{
    private byte[] m_StreamtBuffer = new byte[1024 * 1024];

    private TcpClient m_Socket;
    private int m_BuffIndex = 0;

    private const int m_Size = 4096;

    public SocketOutputStream(TcpClient socket)
    {
        m_Socket = socket;
        m_StreamtBuffer.Initialize();
    }

    //刷新
    public void Flush()
    {
        byte[] tempBuffer = new byte[4096];
        int count = m_StreamtBuffer.Length > m_Size ? m_Size : m_StreamtBuffer.Length;
        for (int i = 0; i < count; i++)
        {
            tempBuffer[i] = (byte)m_StreamtBuffer.GetValue(i);
        }

        int nSize = m_Socket.GetSocket().Send(tempBuffer, count, SocketFlags.None);
        //把m_StreamtBuffer里的数据往前移动，移动开头位置
        Array.Copy(m_StreamtBuffer, nSize, m_StreamtBuffer, 0, m_StreamtBuffer.Length - count);
        m_BuffIndex = m_StreamtBuffer.Length - count;
    }

    public void Write(byte[] data, int length)
    {
        Array.Copy(data, 0, m_StreamtBuffer, m_BuffIndex, length);
        m_BuffIndex += length;
    }
}
