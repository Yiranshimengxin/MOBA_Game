using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ArrayBuffer
{
    private const uint DEFAULT_BUFFER_SIZE = 1024 * 1024;

    private byte[] mBuffer;
    private int mBufferLength;

    public int Length
    {
        get
        {
            return mBufferLength;
        }
    }

    public ArrayBuffer()
    {
        mBuffer = new byte[DEFAULT_BUFFER_SIZE];
    }

    public int WriteBytes(byte[] data, int length)
    {
        if (mBuffer.Length - mBufferLength < length)
        {
            byte[] newBuffer = new byte[mBuffer.Length * 2];
            Array.Copy(mBuffer, 0, newBuffer, 0, mBufferLength);
            mBuffer = newBuffer;

            mBufferLength = mBufferLength + length;
        }
        else
        {
            Array.Copy(data, 0, mBuffer, mBufferLength, length);
            mBufferLength = mBufferLength + length;
        }
        return length;
    }

    public int Read(byte[] data)
    {
        if (data == null)
        {
            return 0;
        }
        int length = 0;
        if (data.Length < mBufferLength)
        {
            Array.Copy(mBuffer, 0, data, 0, data.Length);
            length = data.Length;
        }
        else
        {
            Array.Copy(mBuffer, 0, data, 0, mBufferLength);
            length = mBufferLength;
        }
        return length;
    }

    //归零
    public void Seek(int offest)
    {
        mBufferLength = offest;
    }
}
