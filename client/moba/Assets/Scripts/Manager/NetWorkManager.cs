using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class NetWorkManager : Singleton<NetWorkManager>, ISocketCallBack
{
    private TcpClient m_SocketClient;

    public void Init()
    {
        m_SocketClient = new TcpClient();
        m_SocketClient.Init(this);
        LogManager.Instance.DebugLog("网络管理初始化");

        Connect("127.0.0.1", 10086);
    }

    public void OnConnect()
    {
        Debug.Log("链接服务器成功！");
        Pb.PBLoginReq req = new Pb.PBLoginReq();
        req.Uid = "test";
        Send((int)Pb.MsgID.Login, req);
    }

    public void OnDisconnect()
    {
        LogManager.Instance.ErrorLog("socket disconnect!");
    }

    public void OnReceive(byte[] data, int len)
    {

    }

    public void Connect(string host, int port)
    {
        m_SocketClient.Connect(host, port);
    }

    public void Send(int msgId, Google.Protobuf.IMessage msg)
    {
        Pb.Head head = new Pb.Head();
        head.Id = msgId;
        head.Data = Google.Protobuf.MessageExtensions.ToByteString(msg);

        byte[] msgData = Google.Protobuf.MessageExtensions.ToByteArray(head);
        m_SocketClient.Send(msgData, msgData.Length);
    }

    public void Destroy()
    {

    }

    public void Update()
    {
        m_SocketClient.Update();
    }
}
