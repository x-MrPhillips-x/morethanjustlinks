import { ExclamationCircleIcon } from "@heroicons/react/24/outline";
import Link from "next/link";

export default function loginSuccess(message:string){
    return (
        <div className="flex-1 rounded-lg bg-gray-50 px-6 pb-4 pt-8">
            <div className="round-md bg-gray-50 p-4 md:p-6">
                <ExclamationCircleIcon className="h-5 w-5 text-green-600" />
                <p className="text-sm text-green-600 ml-6">{message}</p>
            </div>
            <div className="mt-6 flex justify-end gap-4">
            <Link
            href="/login"
            className="flex h-10 items-center rounded-lg bg-blue-500 px-4 text-sm font-medium text-white transition-colors hover:bg-green-500"
            >
            Login
            </Link>
            </div>
        </div>
    )
}