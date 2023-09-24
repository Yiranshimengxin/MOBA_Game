using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

//大小端转化
public class BytesHelper : MonoBehaviour
{
    //把一个byte类型数组转换成int类型
    public static int ByteConvertInt32(out int data, byte[] byteData, int byteDataLen, int index, bool LittleEndian = false)
    {
        byte a = byteData[index];
        byte b = byteData[index + 1];
        byte c = byteData[index + 2];
        byte d = byteData[index + 3];
        if (LittleEndian)
        {
            data = (d << 24) | (c << 16) | (b << 8) | a;
        }
        else
        {
            data = (a << 24) | (b << 16) | (c << 8) | d;
        }
        return index + 4;
    }

    //把一个32位整数转成byte[]
    public static int Int32ConvertByte(int data, byte[] byteData, int byteDataLen, int index, bool LittleEndian = false)
    {
        if (index + 4 > byteDataLen)
        {
            return index;
        }
        byte[] netData = BitConverter.GetBytes(data);
        if (LittleEndian)
        {
            byteData[index++] = (byte)(netData[0]);
            byteData[index++] = (byte)(netData[1]);
            byteData[index++] = (byte)(netData[2]);
            byteData[index++] = (byte)(netData[3]);
        }
        else
        {
            byteData[index++] = (byte)(netData[3]);
            byteData[index++] = (byte)(netData[2]);
            byteData[index++] = (byte)(netData[2]);
            byteData[index++] = (byte)(netData[1]);
        }
        return index;
    }
}
