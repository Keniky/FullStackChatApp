import { useNavigate } from 'react-router-dom'
import { SmokeyBackground} from '../components/login-form'
import { useEffect } from 'react'
import { SignInForm } from '@/components/signin-form'

export default function SignInPage(){

    const navigate = useNavigate()

    //once it completed do this
    useEffect(() => {
        fetch("http://localhost:80/api/v1/auth" ,{
                credentials: "include"
            }
        ).then(resp => {
            if (resp.status === 202)  navigate("/chat") 
        })
    }
    , [])


    return(

        <main className="relative w-screen h-screen bg-gray-900">
        <SmokeyBackground className="absolute inset-0" />
        <div className="relative z-10 flex items-center justify-center w-full h-full p-4">
            <SignInForm />
        </div>
        </main>
    )

}