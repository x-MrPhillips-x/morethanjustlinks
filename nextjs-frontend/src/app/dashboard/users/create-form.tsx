'use client'
import Link from "next/link"
import { createUserAccount, State} from "../../lib/actions"
import { Button } from "../../ui/buttons";
import {useFormState} from "react-dom"
import { 
    DevicePhoneMobileIcon, 
    EnvelopeIcon, 
    LockClosedIcon, 
    UserCircleIcon 
} from "@heroicons/react/24/outline";

export default function Form (){

    const initialState:State = {message:null,validationErrors:{}};
    const [state,formAction] = useFormState(createUserAccount,initialState)

    return (
        <form action={formAction}>
            <div className="round-md bg-gray-50 p-4 md:p-6">
                <div className="mb-4">
                    <label htmlFor="name" className="mb-2 block text-sm font-medium">
                        Username
                    </label>
                    <div className="relative">
                        <input className="peer block w-full rounded-md border border-gray-200 py-[9px] pl-10 text-sm outline-2 placeholder:text-gray-500"
                            placeholder="Enter username"
                            name="name"
                            aria-describedby="name-error"
                            minLength={6}
                            id="name"
                            required
                        />
                        <UserCircleIcon className="pointer-events-none absolute left-3 top-1/2 h-[18px] w-[18px] -translate-y-1/2 text-gray-500" />
                    </div>
                    <div id="name-error" aria-live="polite" aria-atomic="true">
                        {state?.validationErrors?.name && 
                            <p className="mt-2 text-sm text-red-500">
                                {state.validationErrors.name}
                            </p>
                        }
                    </div>
                </div>
                <div className="mb-4">
                    <label htmlFor="email" className="mb-2 block text-sm font-medium">
                        Email
                    </label>
                    <div className="relative">
                        <input className="peer block w-full rounded-md border border-gray-200 py-[9px] pl-10 text-sm outline-2 placeholder:text-gray-500"
                            placeholder="Enter email"
                            name="email"
                            aria-describedby="email-error"
                            required
                            id="email"
                        />
                        <EnvelopeIcon className="pointer-events-none absolute left-3 top-1/2 h-[18px] w-[18px] -translate-y-1/2 text-gray-500"/>
                    </div>
                    <div id="email-error" aria-live="polite" aria-atomic="true">
                        {state?.validationErrors?.email && 
                                <p className="mt-2 text-sm text-red-500">
                                    {state.validationErrors.email}
                                </p>
                        }
                    </div>
                </div>
                <div className="mb-4">
                    <label htmlFor="phone" className="mb-2 block text-sm font-medium">
                        Mobile
                    </label>                
                    <div className="relative">
                        <input className="peer block w-full rounded-md border border-gray-200 py-[9px] pl-10 text-sm outline-2 placeholder:text-gray-500"
                            placeholder="Enter mobile phone number"
                            name="phone"
                            required
                            minLength={10}
                            maxLength={10}
                            id="phone"
                            aria-describedby="phone-error"

                        />
                        <DevicePhoneMobileIcon className="pointer-events-none absolute left-3 top-1/2 h-[18px] w-[18px] -translate-y-1/2 text-gray-500"/>
                    </div>
                    <div id="phone-error" aria-live="polite" aria-atomic="true">
                        {state?.validationErrors?.phone && 
                                <p className="mt-2 text-sm text-red-500">
                                    {state.validationErrors.phone}
                                </p>
                        }
                    </div>
                </div>
                <div className="mb-4">
                    <label htmlFor="name" className="mb-2 block text-sm font-medium">
                        Password
                    </label>                
                    <div className="relative">
                        <input className="peer block w-full rounded-md border border-gray-200 py-[9px] pl-10 text-sm outline-2 placeholder:text-gray-500"
                            type="password"
                            name="psword"
                            placeholder="Enter Password"
                            required
                            minLength={6}
                        />
                        <LockClosedIcon className="pointer-events-none absolute left-3 top-1/2 h-[18px] w-[18px] -translate-y-1/2 text-gray-500"/>
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
                <Button type="submit">Create Account</Button>
            </div>

        </form>
    );
}