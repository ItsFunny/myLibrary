/**
*
* @author joker 
* @date 创建时间：2018年4月7日 下午12:21:33
* 
*/
package com.charlie.model;

/**
* 
* @author joker 
* @date 创建时间：2018年4月7日 下午12:21:33
*/
public class TwoTuple<F,S>
{
	public F first;
	public S second;

	public TwoTuple(F first, S second)
	{
		this.first = first;
		this.second = second;
	}

	public F getFirst()
	{
		return first;
	}

	public void setFirst(F first)
	{
		this.first = first;
	}

	public S getSecond()
	{
		return second;
	}

	public void setSecond(S second)
	{
		this.second = second;
	}
}
