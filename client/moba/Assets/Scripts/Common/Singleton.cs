using System.Collections;
using System.Collections.Generic;
using UnityEngine;

//单例模式基类（使用懒汉模式）
public class Singleton<T> where T : class, new()
{
    private static T m_Instance = null;
    public static T Instance
    {
        get
        {
            if(m_Instance == null)
            {
                m_Instance = new T();
            }
            return m_Instance;
        }
    }
}
