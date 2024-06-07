'use server'
import {z} from 'zod';
import { revalidatePath } from 'next/cache';
import { redirect } from 'next/navigation';
import { cookies,headers } from 'next/headers'


const FormSchema = z.object({
    id:z.string(),
    name:z.string({
        invalid_type_error:"Please enter a username"
    }),
    email:z.string({
        invalid_type_error:"Please enter a valid email address"
    }),
    phone:z.string({
        invalid_type_error:"Please enter a valid 10 digit US phone number"
    }),
    psword:z.string({
        invalid_type_error:"Please enter a password"
    }),
    role:z.string(),
})

const CreateUser = FormSchema.omit({id:true,role:true})
const UpdateUser = FormSchema.omit({psword:true})
const DeleteUser = FormSchema.omit({
    name:true,
    email:true,
    phone:true,
    psword:true,
    role:true,
})


export async function getUserSessionCookie(){
    try {
        const responsePromise = await fetch(`${process.env.BACKEND}/incr`)
        const responseCookie = responsePromise.headers.get("set-cookie")
        return responseCookie
    }catch(error){
        return
    }
}


export async function fetchUserById({params}:{params: {id:string}}){
    try {
        const resp = await fetch(`${process.env.BACKEND}/getUser?id=${encodeURIComponent(params.id)}`,{cache:"no-store"})
        const user = Promise.all([resp.json()])
        return user
    } catch(error) {
        return {msg:params.id+" could not be found"}
    }
}

export async function logout(formData:FormData){
    const id = formData.get("id")
    const cookieStore = cookies()
    let cookie = cookieStore.get('set-cookie')

    try {

        await fetch(`${process.env.BACKEND}/logout`,{
            method:'POST',
            body:JSON.stringify({
                "id":id,
            })
        })
        cookies().delete("set-cookie")
        cookie = cookieStore.get('set-cookie')


    } catch(error){
        console.log("something went wrong logging out")
    }

    revalidatePath(process.env.HOST)
    redirect(process.env.HOST)
}

export async function login(prevState:State,formData:FormData){
    const email = formData.get("email")
    const psword = formData.get("psword")

    
    try {
        const initialPromise = await fetch(`${process.env.BACKEND}/login`,{
            method:'POST',
            body:JSON.stringify({
                "email":email,
                "psword":psword,
            }),
        })

        // udpate the headers to pass to NextRequest in middleware
        const updatedHeaders = initialPromise.headers.get('set-cookie')
        cookies().set("set-cookie",updatedHeaders)

        const resp = await Promise.all([initialPromise.json()])
        if (!initialPromise.ok){
            prevState.error = "Unknown email/password combination"
            return prevState
        }
 
    } catch(error) {
        prevState.error = "Something went wrong during authentication"
        return prevState
    }

    revalidatePath(process.env.HOST+"/dashboard")
    redirect(process.env.HOST+"/dashboard")
}

export async function uploadFiles(prevState:State,formData:FormData){
    const file = formData.get("file")
    const id = formData.get("id")

    var data = new FormData()
    data.append("file",file)
    data.append("id",id)

    try {

        const resp = await fetch(`${process.env.BACKEND}/upload`,{
            method:'POST',
            // mode:'cors',
            body:data,
        })

        const uploadStatus = await Promise.all([resp.json()])
        revalidatePath(`/dashboard/users/${id}/upload`);

        prevState.message = uploadStatus[0]["message"] 
        prevState.validationErrors = uploadStatus[0]["errors"] 
        return prevState
    } catch(error) {
        prevState.error = "Something went wrong uploading the file"
        return prevState
    }
}


export async function createUserAccount(prevState:State,formData:FormData){
    const validatedFields = CreateUser.safeParse({
        name:formData.get("name"),
        email:formData.get("email"),
        phone:formData.get("phone"),
        psword:formData.get("psword"),
    })


    if(!validatedFields.success){
        return {
            validationErrors: {},
            message: 'There are some invalid inputs please resolve and try again.'
        }
    }

    const {name,email,phone,psword} = validatedFields.data;

    try {

        const resp = await fetch(`${process.env.BACKEND}/newAccount`,{
            method: 'POST',
            body: JSON.stringify({
                "name":name,
                "email":email,
                "phone":phone,
                "psword":psword,
            }),
        })

        const accountStatus = await Promise.all([resp.json()])
        console.log("response from creating user",accountStatus[0])
        revalidatePath(`/dashboard/users/create`);

        prevState.message = accountStatus[0]["message"] 
        prevState.validationErrors = accountStatus[0]["errors"]
        prevState.error = accountStatus[0]["error"] 
        return prevState
    } catch(error){
        prevState.error = "Something went wrong registering account...";
        return prevState
    }


}

export async function updateUser(prevState:State,formData:FormData) {
    const validatedFields = UpdateUser.safeParse({
        id:formData.get("id"),
        name:formData.get("name"),
        email:formData.get("email"),
        phone:formData.get("phone"),
        role:formData.get("role"),
    })

    if(!validatedFields.success){
        return {
            validationErrors: {},
            message: 'There are some invalid inputs please resolve and try again.'
        }
    }

    const {id,name,phone,email,role} = validatedFields.data;

    try {
        const resp = await fetch(`${process.env.BACKEND}/update`,{
            method: 'POST',
            // mode:'cors',
            body: JSON.stringify({
                "id":id,
                "name":name,
                "email":email,
                "phone":phone,
                "role":role,
            }),
        }) 

        const accountStatus = await Promise.all([resp.json()])

        revalidatePath(`/dashboard/users/${id}/edit`);
        prevState.message = accountStatus[0]["message"]
        prevState.validationErrors = accountStatus[0]["errors"]
        return prevState
    } catch (error){
        // TODO update value because error is undefined
        prevState.validationErrors.name = error;
        return prevState
    }
}

export async function deleteUser(formData:FormData) {
    const id = formData

    try {
        const resp = await fetch(`${process.env.BACKEND}/deleteUser`,{
            method:'POST',
            body:JSON.stringify({
                "id":id,
            })
        })

    } catch(error) {
        return error
    }

    revalidatePath(`/dashboard/users`);
    redirect(`/dashboard/users`);
}

export type State = {
    validationErrors?:{
        name?:string|null;
        email?:string|null;
        phone?:string|null;
    };
    message?:string|null;
    error?:string|null;
};
