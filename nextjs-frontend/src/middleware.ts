import { NextResponse,NextRequest } from "next/server";
import {cookies} from 'next/headers'

// middleware https://nextjs.org/docs/app/building-your-application/routing/middleware
export default async function middleware(request:NextRequest){
    const response = NextResponse.next();
    const cookieStore = cookies()
    let cookie = cookieStore.get('set-cookie')

    // to access /dashboard pages users will need to be logged in with 
    // a session from the backend
    if (cookie && request.nextUrl.pathname.startsWith('/dashboard')){
        response.headers.set('set-cookie',cookie.value)
        return response
    } else if (request.nextUrl.pathname.startsWith('/login')) {
        return response
    }
    else if (request.nextUrl.pathname.startsWith('/create')) {
        return response
    }
    else {
        return NextResponse.rewrite(new URL('/', request.url))   
    }
    
    // return NextResponse.rewrite(new URL('/login', request.url))   
}

export const config = {
    // https://nextjs.org/docs/app/building-your-application/routing/middleware#matcher
    matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
};