import { useEffect } from "react";
import { useRef } from "react";
import { ChatMessageBox } from "./chat-message-box";


export function ChatMessages({chatMessages}){

    const autoScroll = useRef(null)

    useEffect(() => {
        autoScroll.current.scrollIntoView();
    } , [chatMessages]);

    return (
        <div className="flex flex-col gap-[2px] w-full h-full overflow-y-auto pr-2">
            {chatMessages.map((chatMessage) => {
                return (
                    <ChatMessageBox 
                                    message={chatMessage.message}
                                    senderName={chatMessage.name}
                                    senderPfp={chatMessage.pfp}
                                    key={chatMessage.key}/>
                )
            })}

            <div ref={autoScroll}></div>
        </div>
    )

}