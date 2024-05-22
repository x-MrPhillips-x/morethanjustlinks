import { Suspense } from "react";
import { LatestInvoicesSkeleton } from "../../ui/skeletons";
import UsersForAdmin from "../admin-users";
import Breadcrumbs from "./breadcrumbs";


export default function Users() {
    return(
        <main>
      <Breadcrumbs
        breadcrumbs={[
          { label: 'Users', href: '/dashboard/users' },
        ]}
      />
        {/* <div className="mt-6 grid grid-cols-1 gap-6 md:grid-cols-4 lg:grid-cols-8"> */}
        <div className="mt-6">
          <Suspense fallback={<LatestInvoicesSkeleton />}>
            <UsersForAdmin/>
          </Suspense>
        </div>
      </main>
    );
}