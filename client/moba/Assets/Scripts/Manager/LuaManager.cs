using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LuaManager : Singleton<LuaManager>
{
    public void Init()
    {
        LogManager.Instance.DebugLog("Lua管理初始化");
    }

    public void Destroy()
    {

    }
}
