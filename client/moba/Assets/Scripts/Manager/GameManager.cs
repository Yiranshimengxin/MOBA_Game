using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class GameManager:Singleton<GameManager> 
{
    public void Init()
    {
        LogManager.Instance.DebugLog("游戏管理初始化");
    }

    public void Destroy()
    {
        LogManager.Instance.DebugLog("游戏退出销毁");
    }
}
