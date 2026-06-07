import { ArrowRight } from "lucide-react";
import { useNavigate } from "react-router-dom";

export default function SubmitButton({name}){
    const navigate = useNavigate();

    async function logIn() {
        const resp = await fetch("http://localhost:80/api/v1/login", {
            method: "POST",
            credentials: "include",//accept cookies
            headers: {
                'content-type': 'application/json'
            },
            body: JSON.stringify({name:name})
        })
        console.log(resp)
        console.log(resp.ok)
        
        if (resp.ok){
            navigate("/chat")
        }
    }
    return(
        <button
          onClick={logIn}
          type="button"
          className="group w-full flex items-center justify-center py-3 px-4 bg-blue-600 hover:bg-blue-700 rounded-lg text-white font-semibold focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-900 focus:ring-blue-500 transition-all duration-300">
          Sign In
          <ArrowRight className="ml-2 h-5 w-5 transform group-hover:translate-x-1 transition-transform" />
        </button>
    )
}