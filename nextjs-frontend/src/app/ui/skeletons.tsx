// Loading animation
const shimmer =
  'before:absolute before:inset-0 before:-translate-x-full before:animate-[shimmer_2s_infinite] before:bg-gradient-to-r before:from-transparent before:via-white/60 before:to-transparent';

export function InvoiceSkeleton() {
    return (
      <div className="flex flex-row items-center justify-between border-b border-gray-100 py-4">
        <div className="flex items-center">
          <div className="mr-2 h-8 w-8 rounded-full bg-gray-200" />
          <div className="min-w-0">
            <div className="h-5 w-40 rounded-md bg-gray-200" />
            <div className="mt-2 h-4 w-12 rounded-md bg-gray-200" />
          </div>
        </div>
        <div className="mt-2 h-4 w-12 rounded-md bg-gray-200" />
      </div>
    );
  }

export function LatestInvoicesSkeleton() {
    return (
      <div
        className={`${shimmer} relative flex w-full flex-col overflow-hidden md:col-span-4`}
      >
        <div className="mb-4 h-8 w-36 rounded-md bg-gray-100" />
        <div className="flex grow flex-col justify-between rounded-xl bg-gray-100 p-4">
          <div className="bg-white px-6">
            <InvoiceSkeleton />
            <InvoiceSkeleton />
            <InvoiceSkeleton />
            <InvoiceSkeleton />
            <InvoiceSkeleton />
            <div className="flex items-center pb-2 pt-6">
              <div className="h-5 w-5 rounded-full bg-gray-200" />
              <div className="ml-2 h-4 w-20 rounded-md bg-gray-200" />
            </div>
          </div>
        </div>
      </div>
    );
  }