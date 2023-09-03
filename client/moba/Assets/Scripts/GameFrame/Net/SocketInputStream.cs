using System;
using System.Collections;
using System.Collections.Generic;
using System.Net.Sockets;
using UnityEngine;

public class SocketInputStream
{
    private const uint DEFAULT_INPUT_BUFFER_SIZE = 1024 * 256;
    private Array m_InputBuffer;
    private TcpClient m_Socket;

    public SocketInputStream(TcpClient socket)
    {
        m_Socket = socket;
        m_InputBuffer.Initialize();
    }

    //从Socket中读取数据放到m_InputBuffer
    public void Fill()
    {
        byte[] tempData = new byte[4096];
        int datalen= m_Socket.GetSocket().Receive(tempData, 4096, SocketFlags.None);

        Array.Copy(tempData, m_InputBuffer, datalen);
    }
}
