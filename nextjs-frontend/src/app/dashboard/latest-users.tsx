import clsx from "clsx";
import { revalidatePath } from "next/cache";
import { headers } from "next/headers";
import Image from "next/image";
import { redirect } from 'next/navigation';



export async function fetchLatestUsers(){
    // get the cookie-session set at middleware
    const initialHeaders = headers().get('set-cookie')
 
    try {
        const resp = await fetch(`${process.env.BACKEND}/getAllUsers`,{ 
            cache:'no-store',
            headers: {
                cookie:initialHeaders,
            }        
            // const initialPromise = await fetch(`${process.env.BACKEND}/incr`,{
            //     headers:{
            //         cookie:initialHeaders
            //     },
            // })
        })
        const users = await Promise.all([resp.json()])
        return users
    } catch (error) {
        return [{"error":"something went wrong..."}]

    }

}

export default async function LatestUsers(){
    const latestUsers = await fetchLatestUsers();
    
    return (
        <div className="flex w-full flex-col md:col-span-4">
            <h2 className="mb-4 text-xl md:tex-2xl">Latest Users</h2>
            <div className="flex grow flex-col justify-between rounded-xl bg-gray-50 p-4">
                <div className="bg-white px-6">
                    {latestUsers[0]["message"]?.map((user,i)=> {
                        return(
                            <div key={user.ID}
                                className={clsx(
                                    "flex flex-row items-center justify-between py-4",
                                    {
                                        "border-t": i !==0,
                                    },
                                )}
                            >
                                <div className="flex items-center">
                                    <Image 
                                        className="mr-4 rounded-full"
                                        src={user.profilePic} 
                                        alt={user.name}
                                        width={32}
                                        height={32}
                                    />
                                    <div className="min-w-0">
                                        <p className="truncate text-sm font-semibold md:text-base">
                                            {user.name}
                                        </p>
                                        <p className="hidden text-sm text-gray-500 sm:block">
                                            {user.email}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        );
                    })}
                </div>
            </div>
        </div>
    );
}