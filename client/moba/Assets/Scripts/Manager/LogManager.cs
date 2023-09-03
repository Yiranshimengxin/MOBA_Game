using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LogManager : Singleton<LogManager>
{
    private LogSystem m_Log;
    public void Init()
    {
        m_Log = new LogSystem();
        Debug.Log("游戏日志初始化");
        m_Log.OpenFile("");
        m_Log.SetLevel(0);
    }

    public void Destroy()
    {

    }

    //debug 信息
    public void DebugLog(string text)
    {
        //Debug.Log(text);
        m_Log.Debugf("这是一个Debug日志");
        m_Log.Infof("这是一个Info日志");
        m_Log.Warningf("这是一个Warning日志");
        m_Log.Errorf("这是一个Error日志");
    }

    //warning 警告
    public void WarningLog(string text)
    {
        //Debug.Log(text);
    }

    //error 报错
    public void ErrorLog(string text)
    {
        m_Log.Errorf(text);
    }
}
