import { ChatMembers } from "@/components/room/chat-members";
import { MainChatArea } from "@/components/room/main-chat-area";
import { RoomHeader } from "@/components/room/room-header";
import { RoomMessagesAndMemmbersContext } from "@/context/roomContext";
import { useState } from "react";



export default function ChatRoomPage(){


    const [roomMembers , setRoomMembers] =  useState([])
    const [chatMessages , setChatMessages] = useState([])



    return (
        <RoomMessagesAndMemmbersContext.Provider value={{roomMembers , setRoomMembers , chatMessages , setChatMessages}}>
            <RoomHeader />
            <div className="flex flex-row h-[95vh]">
                <MainChatArea />            
                <ChatMembers />
            </div>
        </RoomMessagesAndMemmbersContext.Provider>
        

    )
}