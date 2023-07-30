using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ResourceManager : Singleton<ResourceManager>
{
    public void Init()
    {
        LogManager.Instance.DebugLog("资源管理初始化");
    }

    public void Destroy()
    {

    }
}
