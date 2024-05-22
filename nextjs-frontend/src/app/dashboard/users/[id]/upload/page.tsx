import { fetchUserById } from "../../../../lib/actions";
import Breadcrumbs from "../../breadcrumbs";
import Form from "../../upload-form";

export  default async function Page({ params }: { params: { id: string } }) {
    const id = params.id;
    const user = await fetchUserById({params:{id:id}})
    return(
        <main>
        <Breadcrumbs
          breadcrumbs={[
            { label: 'Users', href: '/dashboard/users' },
            {
              label: `Upload files for ${user[0]['message'].name}`,
              href: `/dashboard/users/${id}/upload`,
              active: true,
            },
          ]}
        />
        <Form user={user}></Form>
      </main>
    );
}