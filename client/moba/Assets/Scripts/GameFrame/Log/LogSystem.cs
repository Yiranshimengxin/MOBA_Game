using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEngine;

public class LogSystem : MonoBehaviour
{
    private enum LogLevel
    {
        Debug = 0,
        Info = 1,
        Warning = 2,
        Error = 3
    }

    //游戏日志文件存放的目录
    private string m_LogFilePath = "";

    //日志文件写入流
    private FileStream m_LogFile = null;

    //文件写入器
    private BinaryWriter m_Write = null;

    //记录当前文件的大小
    private int m_FileSize = 0;

    //当前输出的log有等级
    private int m_LogLevel = (int)LogLevel.Debug;


    //创建日志文件
    public bool OpenFile(string filePath)
    {
        m_LogFilePath = filePath;
        try
        {
            if (m_LogFilePath == "")
            {
                //当前游戏运行的平台是苹果
                if (Application.platform == RuntimePlatform.IPhonePlayer)
                {

                }
                //当前游戏运行的平台是安卓
                else if (Application.platform == RuntimePlatform.Android)
                {

                }
                //当前游戏运行的平台为Windows系统
                else if (Application.platform == RuntimePlatform.WindowsPlayer)
                {

                }
                //当前游戏运行的平台为Windows或苹果编辑器模式（开发模式）
                else if (Application.platform == RuntimePlatform.WindowsEditor || Application.platform == RuntimePlatform.OSXEditor)
                {
                    string path = Application.dataPath;
                    path = path.Substring(0, path.IndexOf("Asset"));
                    m_LogFilePath = Path.Combine(path, string.Format("{0}_ClientLog.txt", System.DateTime.Now.ToString("yyyy_MM_dd_HH_mm_ss")));
                }
            }

            //如果有过期的日志文件存在，则删除

            //获取当前文件所在目录
            string strDir = Path.GetDirectoryName(m_LogFilePath);
            if (Directory.Exists(strDir))
            {
                //获取文件夹下的所有txt文件，循环删除
                string[] oldLogFiles = Directory.GetFiles(strDir, "*.txt");
                foreach (string fileName in oldLogFiles)
                {
                    //删除文件
                    File.SetAttributes(fileName, FileAttributes.Normal);
                    File.Delete(fileName);
                }
            }

            m_LogFile = new FileStream(m_LogFilePath, FileMode.Create, FileAccess.ReadWrite, FileShare.Delete | FileShare.Read);
            /*创建日志文件流：m_LogFile = new FileStream(m_LogFilePath, FileMode.Create, FileAccess.ReadWrite, FileShare.Delete | FileShare.Read);
            使用确定的日志文件路径m_LogFilePath创建一个新的文件流m_LogFile。
            文件流的访问模式是FileMode.Create，这将创建一个新的日志文件或覆盖已存在的同名文件。
            文件流的访问权限是FileAccess.ReadWrite，允许读取和写入文件。
            文件共享模式是FileShare.Delete | FileShare.Read，允许其他进程删除文件，但允许读取文件的访问。*/

            m_Write = new BinaryWriter(m_LogFile);

        }
        catch (Exception e)
        {
            Debug.Log(e.ToString());
            return false;
        }
        return true;
    }

    //关闭文件日志系统
    public void CloseFile()
    {
        if (m_Write != null)
        {
            m_Write.Close();
            m_Write = null;
        }
        if (m_LogFile != null)
        {
            m_LogFile.Close();
            m_LogFile.Dispose();
            m_LogFile = null;
        }
    }

    private void PrintLog(string msg)
    {
        string timeStr = DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss:ffff");
        timeStr += " ";
        timeStr += msg;
        timeStr += "\n";

        m_Write.Write(timeStr);
        m_Write.Flush();  //刷新到缓冲区

        m_FileSize += timeStr.Length;
        //如果当前的日志文件大于1K，则新建一个文件继续写入日志
        if (m_FileSize >= 1024 * 1024 * 1024)
        {
            //新建一个文件
            NewLogFile();
        }
    }

    private void NewLogFile()
    {
        if (m_LogFile != null)
        {
            m_LogFile.Close();
            m_LogFile.Dispose();
            m_LogFile = null;
        }
        if (m_Write != null)
        {
            m_Write.Close();
            m_Write = null;
        }
        //当前游戏运行的平台是苹果
        if (Application.platform == RuntimePlatform.IPhonePlayer)
        {

        }
        //当前游戏运行的平台是安卓
        else if (Application.platform == RuntimePlatform.Android)
        {

        }
        //当前游戏运行的平台为Windows系统
        else if (Application.platform == RuntimePlatform.WindowsPlayer)
        {

        }
        //当前游戏运行的平台为Windows或苹果编辑器模式（开发模式）
        else if (Application.platform == RuntimePlatform.WindowsEditor || Application.platform == RuntimePlatform.OSXEditor)
        {
            string path = Application.dataPath;
            path = path.Substring(0, path.IndexOf("Asset"));
            m_LogFilePath = Path.Combine(path, string.Format("{0}_ClientLog.txt", System.DateTime.Now.ToString("yyyy_MM_dd_HH_mm_ss")));
        }

        m_LogFile = new FileStream(m_LogFilePath, FileMode.Create, FileAccess.ReadWrite, FileShare.Delete | FileShare.Read);
        m_Write = new BinaryWriter(m_LogFile);
    }

    public void Debugf(string format)
    {
        if (m_LogLevel > (int)LogLevel.Debug)
        {
            return;
        }
        string strDebug = "[Debug]";
        strDebug += format;
        PrintLog(strDebug);
    }

    public void Infof(string format)
    {
        if (m_LogLevel > (int)LogLevel.Info)
        {
            return;
        }
        string strInfo = "[Info]";
        strInfo += format;
        PrintLog(strInfo);
    }

    public void Warningf(string format)
    {
        if (m_LogLevel > (int)LogLevel.Warning)
        {
            return;
        }
        string strWarning = "[Warning]";
        strWarning += format;
        PrintLog(strWarning);
    }

    public void Errorf(string format)
    {
        if (m_LogLevel > (int)LogLevel.Error)
        {
            return;
        }
        string strError = "[Error]";
        strError += format;
        PrintLog(strError);
    }

    public void SetLevel(int level)
    {
        m_LogLevel = level;
    }
}
