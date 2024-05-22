'use client'

import Link from 'next/link';
import { PencilIcon, TrashIcon} from '@heroicons/react/24/outline';
import { deleteUser } from '../lib/actions';
import {useFormStatus} from 'react-dom';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
}

export function Button({children,className, ...rest}: ButtonProps){
  const {pending} = useFormStatus();
  return (
    <button 
      {...rest}
      className="flex h-10 items-center rounded-lg bg-blue-500 px-4 text-sm font-medium text-white transition-colors hover:bg-green-500 aria-disabled:cursor-not-allowed aria-disabled:opacity-50 disabled:bg-gray-500 disabled:text-white"
      disabled={pending}
      aria-disabled={pending}
    > 
      {children}
    </button>
  );
}

export function UpdateUser({ id }: { id: string }) {
    return (
      <Link
        href={`/dashboard/users/${id}/edit`}
        className="rounded-md border p-2 hover:bg-gray-100"
      >
        <PencilIcon className="w-5" />
      </Link>
    );
}

export function DeleteUser({ id }: { id: string }) {
  const deleteUserWithID = deleteUser.bind(null,id)
    return (
      <form action={deleteUserWithID}>
      <button className="rounded-md border p-2 text-red-500 hover:bg-gray-100">
        <span className="sr-only">Delete</span>
        <TrashIcon className="w-5" />
      </button>
    </form>
    );
}