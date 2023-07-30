using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PatchManager : Singleton<PatchManager>
{
    public void Init()
    {
        LogManager.Instance.DebugLog("资源更新始化");
    }

    public void Destroy()
    {

    }
}
