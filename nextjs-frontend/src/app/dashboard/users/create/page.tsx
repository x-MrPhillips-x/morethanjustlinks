import Breadcrumbs from "../breadcrumbs";
import Form from "../create-form";

export default async function Page() {
    // fetch customers
    return (
        <main>
            <Breadcrumbs
                breadcrumbs={[
                    {label:'Users', href:'/dashboard/users'},
                    {label:'Create Users', href:'/dashboard/users/create', active:true},
                ]}
            />
            <Form/>
        </main>
    );
}