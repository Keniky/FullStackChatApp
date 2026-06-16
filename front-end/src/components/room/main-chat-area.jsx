import { ChatMessages } from "./chat-messages";
import { useContext } from "react";
import { RoomContext, RoomMessagesAndMemmbersContext } from "@/context/roomContext";
import { ChatInput } from "./chat-input";
import { useEffect } from "react";
import { useRef } from "react";

export function MainChatArea(){

    const messagesWebSocket = useRef(null)
    const {chatMessages, setChatMessages} = useContext(RoomMessagesAndMemmbersContext)
    const {roomId} = useContext(RoomContext)
    //join room with cookie and room id
    useEffect(() =>{
        // const url = new URL("ws://localhost:80/api/v1/chat")
        // url.searchParams.set("room_id", roomId)
        // const ws = new WebSocket(url);
        // setWebSocket(ws);
        // ws.onmessage = (msg) => console.log(msg.data);
        // ws.onopen = () => ws.send(JSON.stringify({msg: "bruh"}));

        const messagesUrl = new URL("ws://localhost:80/api/v1/chat");
        messagesUrl.searchParams.set("room_id" , roomId);
        messagesWebSocket.current = new WebSocket(messagesUrl);

        messagesWebSocket.current.onmessage = (msg) => {
            const data = JSON.parse(msg.data);
            setChatMessages(prev => [...prev , {
                                                message:data.message,
                                                name:data.name,
                                                pfp:data.pfp,
                                                key:crypto.randomUUID()}])
        } 

    }, [])

    return(
        <div className="flex flex-col justify-between w-full h-[95vh] bg-gradient-to-b from-[#313338] via-gray-[#202124] to-[#313338] ">
            <div className="flex-1 overflow-scroll">
                <ChatMessages chatMessages={chatMessages}/> 
            </div>
                <ChatInput messagesWebSocket={messagesWebSocket}/>
        </div>
    )
}