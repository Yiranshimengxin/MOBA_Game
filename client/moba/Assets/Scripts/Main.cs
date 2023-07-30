using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Main : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        LogManager.Instance.Init();
        GameManager.Instance.Init();
        NetWorkManager.Instance.Init();
        ResourceManager.Instance.Init();
        SoundManager.Instance.Init();
        VideoManager.Instance.Init();
        PatchManager.Instance.Init();
        LuaManager.Instance.Init();
    }

    // Update is called once per frame
    void Update()
    {

    }

    private void FixedUpdate()
    {
        
    }

    private void LateUpdate()
    {
        
    }


    public void OnApplicationQuit()
    {
        LogManager.Instance.Destroy();
    }
}
