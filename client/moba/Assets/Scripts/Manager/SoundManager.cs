using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SoundManager : Singleton<SoundManager>
{
    public void Init()
    {
        LogManager.Instance.DebugLog("声音管理初始化");
    }

    public void Destroy()
    {

    }
}
