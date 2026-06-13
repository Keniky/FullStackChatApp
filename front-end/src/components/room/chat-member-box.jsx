

export function ChatMemberBox({roomMember}){

    console.log("test")
    return(
        <div className="mb-4 hover:bg-gray-500 ">
            {roomMember.name}
        </div>
    )
}