import Link from 'next/link';

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col p-6">
      <div className="flex h-20 shrink-0 items-end rounded-lg bg-gradient-to-l from-blue-500 p-4 md:h-52">
          <h2 className="text-5xl text-blue-500 font-bold">&gt;&gt;&gt;JustLinks🔗</h2> 
      </div>
      <div className="mt-4 flex grow flex-col gap-4 md:flex-row">
        <div className="flex flex-col justify-center gap-6 rounded-lg bg-gray-50 px-6 py-10 md:w-2/5 md:px-20">
          <p className="text-xl text-gray-800 md:text-3xl md:leading-normal">
            <strong>Welcome to More Than Just Links!!!</strong>
          </p>
          <p>This is an example application created by <a href="https://github.com/x-MrPhillips-x" className="underline">MrPhillips</a> to demonstrate
            my full stack application:
          </p>
          <ul>
            <li>Docker Containerization</li>
            <li>Golang + Gin</li>
            <li>Sql/Postgres</li>
            <li>Next.js</li>
            <li>🚧 React Native</li>
          </ul>
          <Link
            href="/login"
            className="flex items-center gap-5 self-start rounded-lg bg-blue-500 px-6 py-3 text-sm font-medium text-white transition-colors hover:bg-blue-400 md:text-base"
          >
            <span>Log In</span>
          </Link>
          <Link
            href="/create"
            className="flex items-center gap-5 self-start rounded-lg bg-blue-500 px-6 py-3 text-sm font-medium text-white transition-colors hover:bg-blue-400 md:text-base"
          >
            <span>Create Account</span>
          </Link>
        </div>
        <div className="flex items-center justify-center p-6 md:w-3/5 md:px-28 md:py-12">
        🚧  Add Hero Images Here
        </div>
      </div>
    </main>
  );
}
