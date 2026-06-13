import { RoomContext } from "@/context/roomContext";
import { useState } from "react";
import { useContext } from "react";

export function ChatInput({messagesWebSocket}){

    const [message , setMessage] = useState('');
    const {name } = useContext(RoomContext)

    const registerInput = (event) => {
        setMessage(event.target.value);
    }
    const keyEvent = (event) => {
        if(event.key == "Enter"){
            sendMessage()
        }
    }
    const sendMessage = () => {
        //sendMessage
        //replace this with send message to server
        messagesWebSocket.current.send(JSON.stringify({message:message , name:name})) 
        setMessage('')
    }
    
    return (
        <div className="flex flex-row  items-center gap-4 pb-6 pt-2 flex  justify-center">
            <input 
                type="text" 
                placeholder = "Send a message..."
                onChange = {registerInput}
                onKeyDown = {keyEvent}
                value = {message}
                className=" max-w-2xl flex-1 bg-white/90 text-black rounded-xl py-2 px-4 max-w-2xl shadow-inner placeholder-gray-500"
            />    
            <button type="button" 
            className="bg-gradient-to-r from-indigo-500 to-purple-500 text-white font-semibold px-5 py-2 rounded-xl shadow-md hover:from-indigo-600 hover:to-purple-600 transition" 
            onClick = {sendMessage}>
                Send
            </button>
        </div>
    );
}