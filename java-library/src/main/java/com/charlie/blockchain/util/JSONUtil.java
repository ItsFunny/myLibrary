package com.charlie.blockchain.util;


import com.alibaba.fastjson.JSON;
import com.google.gson.*;

import java.lang.reflect.Type;
import java.util.List;
import java.util.Map;

/**
 * @author joker
 * @date 创建时间：2018年5月14日 上午10:54:40
 */
public class JSONUtil
{
    private static Gson gson = null;

    static
    {
        gson=new Gson();
//        GsonBuilder gsonBuilder = new GsonBuilder();
//        gsonBuilder.setDateFormat("yyyy-MM-dd HH:mm:ss");
//        gsonBuilder.registerTypeAdapter(Serializable.class, new GsonInterfaceAdapter());
//        gsonBuilder.registerTypeAdapter(Date.class, new JsonDeserializer<Date>()
//        {
//            private SimpleDateFormat dtf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
//
//            @Override
//            public Date deserialize(JsonElement json, Type typeOfT, JsonDeserializationContext context)
//                    throws JsonParseException
//            {
//                try
//                {
//                    return new Date(json.getAsJsonPrimitive().getAsLong());
//
//                } catch (NumberFormatException e)
//                {
//                    // Get the json element as a String and parse it to get a Date
//                    try
//                    {
//                        return dtf.parse(json.getAsString());
//                    } catch (ParseException e2)
//                    {
//                        // Throw a JsonParseException in case of a parsing error
//                        throw new JsonParseException(e);
//                    }
//                }
//
//            }
//        });
//        gsonBuilder.setLongSerializationPolicy(LongSerializationPolicy.STRING);
////		gsonBuilder.registerTypeAdapter(Date.class, new JsonSerializer<Date>()
////		{
////
////			@Override
////			public void serialize(Date value, JsonGenerator gen, SerializerProvider serializers) throws IOException
////			{
////
////			}
////
////		});
//        gson = gsonBuilder.create();
    }

    public static <K, V> Map<K, V> json2Map(String json, Type type)
    {
        return gson.fromJson(json, type);
    }

    public static void main(String[] args)
    {
        String[] arr = new String[]
                {"1", "2"};
        System.out.println(arr);
    }

    public static String obj2Json(Object object)
    {
        String json = gson.toJson(object);

        return json;
    }

    public static <T> List<T> json2List(String json, Type type)
    {
        return gson.fromJson(json, type);
    }

//    public static <T> T json2Object(String json, Class<T> c)
//    {
//        T t = gson.fromJson(json, c);
//        return t;
//    }

    public static <T> T json2Object(String json, Type c)
    {
        return JSON.parseObject(json, c);
    }
}
