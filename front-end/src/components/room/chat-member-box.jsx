

export function ChatMemberBox({roomMember}){


    return(
        <div className="mb-4 hover:bg-gray-500 ">
            <div className="w-90 h-90 rounded-full border border-black border-solid">
                {roomMember.pfp && <img src="/favicon.png" alt="./static/app/public/a.png"/>}
            </div>
            {roomMember.name}
        </div>
    )
}