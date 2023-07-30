using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class NetWorkManager : Singleton<NetWorkManager>
{
    public void Init()
    {
        LogManager.Instance.DebugLog("网络管理初始化");
    }

    public void Destroy()
    {

    }
}
