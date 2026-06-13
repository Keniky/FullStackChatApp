import {  RoomContext, RoomMessagesAndMemmbersContext } from "@/context/roomContext"
import { useContext } from "react"
import { ChatMemberBox } from "./chat-member-box"
import { useEffect } from "react"
import { useRef } from "react"

export function ChatMembers(){

    const membersWebSocket = useRef(null)
    const {roomId} = useContext(RoomContext)
    const {roomMembers , setRoomMembers} = useContext(RoomMessagesAndMemmbersContext)

    //join room with cookie and room id
    useEffect(() =>{
        // const url = new URL("ws://localhost:80/api/v1/chat")
        // url.searchParams.set("room_id", roomId)
        // const ws = new WebSocket(url);
        // setWebSocket(ws);
        // ws.onmessage = (msg) => console.log(msg.data);
        // ws.onopen = () => ws.send(JSON.stringify({msg: "bruh"}));

        const membersUrl = new URL("ws://localhost:80/api/v1/chat/members");
        membersUrl.searchParams.set("room_id" , roomId);
        membersWebSocket.current = new WebSocket(membersUrl);
        membersWebSocket.current.onmessage = (msg) =>{
            const data = JSON.parse(msg.data);

            setRoomMembers(data.members)
        } 

    }, [])

    return (
        <div className="h-full w-[16rem] b-l b-black b-solid flex flex-col text-white border-l border-black
                        pl-4 bg-gradient-to-b from-[#313338] via-gray-[#202124] to-[#313338] ">
            <span className="mb-9 border-b border-black pb-4">Chat Members Are </span>
            {
                roomMembers.map((roomMember) => {
                    return (<ChatMemberBox roomMember={roomMember}/>

                    )})}
        </div>
    )
}