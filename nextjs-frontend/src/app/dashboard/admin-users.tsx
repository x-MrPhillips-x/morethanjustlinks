import Image from "next/image";
import { fetchLatestUsers } from "./latest-users";
import { DeleteUser, UpdateUser } from "../ui/buttons";


export default async function UsersForAdmin(){
    const users = await fetchLatestUsers();

    return (
        <div className="mt-6 flow-root">
        <div className="inline-block min-w-full align-middle">
          <div className="rounded-lg bg-gray-50 p-2 md:pt-0">
            <div className="md:hidden">
              {users[0]["message"]?.map((user) => (
                <div
                  key={user.ID}
                  className="mb-2 w-full rounded-md bg-white p-4"
                >
                  <div className="flex items-center justify-between border-b pb-4">
                    <div>
                      <div className="mb-2 flex items-center">
                        <Image
                          src={user.profilePic}
                          className="mr-2 rounded-full"
                          width={28}
                          height={28}
                          alt={`${user.name}'s profile picture`}
                        />
                        <p>{user.name}</p>
                      </div>
                      <p className="text-sm text-gray-500">{user.role}</p>
                    </div>
                    {user.ID}
                    {/* <InvoiceStatus status={invoice.status} /> */}
                  </div>
                  <div className="flex w-full items-center justify-between pt-4">
                    <div>
                      <p className="text-xl font-medium">
                        {user.email}
                        {/* {formatCurrency(invoice.amount)} */}
                      </p>
                      <p>
                        {user.phone}
                        {/* {formatDateToLocal(invoice.date)} */}
                      </p>
                    </div>
                    <div className="flex justify-end gap-2">
                      <UpdateUser id={user.ID} />
                      <DeleteUser id={user.ID} />
                    </div>
                  </div>
                </div>
              ))}
            </div>
            <table className="hidden min-w-full text-gray-900 md:table">
              <thead className="rounded-lg text-left text-sm font-normal">
                <tr>
                  <th scope="col" className="px-4 py-5 font-medium sm:pl-6">
                    User
                  </th>
                  <th scope="col" className="px-3 py-5 font-medium">
                    UUID
                  </th>
                  <th scope="col" className="px-3 py-5 font-medium">
                    Email
                  </th>
                  <th scope="col" className="px-3 py-5 font-medium">
                    Phone
                  </th>
                  <th scope="col" className="px-3 py-5 font-medium">
                    Role
                  </th>
                  <th scope="col" className="relative py-3 pl-6 pr-3">
                    <span className="sr-only">Edit</span>
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white">
                {users[0]["message"]?.map((user) => (
                  <tr
                    key={user.ID}
                    className="w-full border-b py-3 text-sm last-of-type:border-none [&:first-child>td:first-child]:rounded-tl-lg [&:first-child>td:last-child]:rounded-tr-lg [&:last-child>td:first-child]:rounded-bl-lg [&:last-child>td:last-child]:rounded-br-lg"
                  >
                    <td className="whitespace-nowrap py-3 pl-6 pr-3">
                      <div className="flex items-center gap-3">

                        <Image
                          src={user.profilePic}
                          className="rounded-full"
                          width={28}
                          height={28}
                          alt={`${user.name}'s profile picture`}
                        />
                        <p>{user.name}</p>
                      </div>
                    </td>
                    <td className="whitespace-nowrap px-3 py-3">
                      {user.ID}
                    </td>
                    <td className="whitespace-nowrap px-3 py-3">
                        {user.email}

                      {/* {formatCurrency(invoice.amount)} */}
                    </td>
                    <td className="whitespace-nowrap px-3 py-3">
                    {user.phone}

                      {/* {formatDateToLocal(invoice.date)} */}
                    </td>
                    <td className="whitespace-nowrap px-3 py-3">
                        {user.role}
                      {/* <InvoiceStatus status={invoice.status} /> */}
                    </td>
                    <td className="whitespace-nowrap py-3 pl-6 pr-3">
                      <div className="flex justify-end gap-3">
                        <UpdateUser id={user.ID} />
                        <DeleteUser id={user.ID} />
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    );
}