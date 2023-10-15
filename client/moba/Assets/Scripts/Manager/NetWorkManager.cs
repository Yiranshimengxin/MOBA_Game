using Pb;
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class NetWorkManager : ISocketCallBack
{
    private static NetWorkManager mInstance = null;

    public static NetWorkManager Instance
    {
        get
        {
            if (mInstance == null)
            {
                mInstance = new NetWorkManager();
            }
            return mInstance;
        }
    }

    private const int MSG_HEAD_SIZE = 4;

    private TcpClient m_SocketClient;
    private byte[] mReceiveBuffer = new byte[1024 * 1024];
    private int mReceiveLength = 0;

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
        Send((int)Pb.MsgID.LoginReq, req);
    }

    public void OnDisconnect()
    {
        LogManager.Instance.ErrorLog("socket disconnect!");
    }

    public void OnReceive(byte[] data, int len)
    {
        Debug.Log("OnReceive len " + len);
        Array.Copy(data, 0, mReceiveBuffer, mReceiveLength, len);
        mReceiveLength = mReceiveLength + len;
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

        ProcessMessage();
    }

    private void ProcessMessage()
    {
        int readIndex = 0;
        do
        {
            if (mReceiveLength < MSG_HEAD_SIZE)
            {
                break;
            }
            int size = 0;
            BytesHelper.ByteConvertInt32(out size, mReceiveBuffer, 4, readIndex, false);
            if (mReceiveLength - 4 < size)
            {
                break;
            }

            mReceiveLength = mReceiveLength - 4;
            readIndex += 4;

            byte[] bytes = new byte[size];
            Array.Copy(mReceiveBuffer, readIndex, bytes, 0, size);
            readIndex += size;
            mReceiveLength = mReceiveLength - size;

            OnMessage(bytes, size);
        } while (true);
    }

    private void OnMessage(byte[] data, int length)
    {
        Head head = Head.Parser.ParseFrom(data, 0, length);

        MsgID id = (MsgID)head.Id;
        Debug.Log("receive msg id =" + id);

        //switch (id)
        //{
        //    case MsgID.LoginRsp:
        //        LoginSys.Instance.LoginRsp(Pb.PBLoginRsp.Parser.ParseFrom(head.Data));
        //        break;
        //    case MsgID.MatchRsp:
        //        LobbySys.Instance.MatchRsp(Pb.PbMatchRsp.Parser.ParseFrom(head.Data));
        //        break;
        //    case MsgID.NotifyConfirm:
        //        LobbySys.Instance.NotifyConfirm2(Pb.PBnotifyConfirm.Parser.ParseFrom(head.Data));
        //        break;
        //    case MsgID.NotifySelect:
        //        LobbySys.Instance.NotifySelect();
        //        break;
        //    case MsgID.NotifyLoad:
        //        LobbySys.Instance.NotifyLoadRes2(Pb.PBNotifyLoad.Parser.ParseFrom(head.Data));
        //        break;
        //    case MsgID.NotifyProgress:
        //        BattleSys.Instance.NotifyLoadPrg();
        //        break;
        //    case MsgID.NotifyBattleStart:
        //        BattleSys.Instance.RspBattleStart2(Pb.PBNotifyBattleStart.Parser.ParseFrom(head.Data));
        //        break;
        //    case MsgID.NotifyOperate:
        //        BattleSys.Instance.NotifyOpKey2(Pb.PBNotifyOperate.Parser.ParseFrom(head.Data));
        //        break;
        //    default:
        //        Debug.Log("error msg id");
        //        break;
        //}
    }
}
