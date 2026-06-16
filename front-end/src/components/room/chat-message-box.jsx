import { RoomContext } from "@/context/roomContext"
import { useContext } from "react"
import dayjs from "dayjs"

export function ChatMessageBox({message , senderName, senderPfp}){
    const {name} = useContext(RoomContext)
    //image then {information under it message it self }
    //add images later
    return (
        <div className={`flex flex-1 ${senderName == name ? "flex-row-reverse" : "flex-row"}  items-center  m-[1rem]  `}> 
            <div className="h-90 w-90 border border-black rounded-full border-solid">
                <img src={senderPfp} alt="./static/app/public/a.png" ></img>
            </div>
            <div className="flex flex-col px-[2rem] max-w-[30rem]">
                <section className="flex flex-row items-center gap-[1rem] flex justify-between pl-[10px] ">
                    <span className="text-[18px] font-medium">{senderName}</span>     
                    <span className="text-[12px] text-gray-300">{dayjs().format('HH:mm:ss')}</span>
                </section>
                <section className="rounded-2xl p-4 bg-gradient-to-r from-gray-700 via-gray-600 to-gray-700 text-white shadow-lg h-auto overflow-visible break-words">
                    {message}
                </section>
            </div> 
        </div>
    )
}