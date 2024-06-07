'use client'
import { State, uploadFiles } from "../../lib/actions";
import {useFormState} from "react-dom"
import Link from "next/link";
import { Button } from "../../ui/buttons";

export default function Form({user}){
    const initialState:State = {message:null,validationErrors:{}};
    const [state,formAction] = useFormState(uploadFiles,initialState)
    
    return(
        <form action={formAction}>
            <div className="round-md bg-gray-50 p-4 md:p-6">
            <input className="peer block w-full rounded-md border border-gray-200 py-[9px] pl-10 text-sm outline-2 placeholder:text-gray-500"
                        defaultValue={user[0]['message'].ID}
                        name="id"
                        type="hidden"
                /> 
                <div className="mb-4">
                    <label htmlFor="name" className="mb-2 block text-sm font-medium">
                        File Upload
                    </label>
                    <div className="relative">
                        <input 
                            name="file"
                            type="file" 
                            className="block w-full text-sm border border-gray-200 rounded-lg cursor-pointer focus:outline-none" 
                            id="file_input"/>
                    </div>
                </div>
            </div>
            <div className="mt-6 flex justify-end gap-4">
                <Link
                href="/dashboard/users"
                className="flex h-10 items-center rounded-lg bg-gray-100 px-4 text-sm font-medium text-gray-600 transition-colors hover:bg-gray-200"
                >
                Cancel
                </Link>
                <Button type="submit">Upload File</Button>
            </div>
        </form>
    );
}