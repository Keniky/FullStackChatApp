

export function ChatMemberBox({roomMember}){


    return(
        <div className="mb-4 hover:bg-gray-500 flex flex-row">
            <div className="w-45 h-45 rounded-full border border-black border-solid">
                {roomMember.pfp && <img src={"/" + roomMember.pfp} onError={(e) => {e.currentTarget.src = '/favicon.png'}} className="w-12 rounder-full "/>}
            </div>
            {roomMember.name}
        </div>
    )
}