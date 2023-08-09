<script>
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    
    let client_ipaddress = "";
    let client_timezone = "";
    const schema = yup.object().shape({
        username: yup.string().required().matches(/^[a-zA-z0-9]+$/, "Username must Character A-Z or a-z or 1-9 ").min(3).max(30),
        password: yup.string().required().min(4).max(30)
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            username: "",
            password: ""
        },
        validationSchema: schema,
        onSubmit:(values) => {
            handleSave(values.username,values.password)
        }
    })
    async function handleSave(username,password) {
        const res = await fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username: username,
                password: password,
                ipaddress: client_ipaddress,
                timezone: client_timezone,
            }),
        });
        const json = await res.json();
        if (json.status == 400) {
            alert(json.message);
            $form.username = ""
            $form.password = ""
        } else {
            localStorage.setItem("token", json.token);
            window.location.href = "/";
        }
    }
    async function initTimezone() {
        const res = await fetch("https://ipinfo.io/json?token=0d10fdc946df5a");
        if(res.status == 200){
            if (!res.ok) {
                const message = `An error has occured: ${res.status}`;
                throw new Error(message);
            } else {
                const json = await res.json();
                client_ipaddress = json.ip;
                client_timezone = json.timezone;
            }
        }else{
            client_ipaddress = "111.111.111.111";
            client_timezone = "Asia/Jakarta";
        }
        
    }
    initTimezone();
    $:{
        if ($errors.username || $errors.password){
            alert($errors.username+"\n"+$errors.password)
            $form.username = ""
            $form.password = ""
        }
    }
</script>
<div class="container" style="margin-top: 20px;">
    <div class="row">
        <div class="col-md-4 offset-md-4">
            <div class="card mt-1" style="">
                <div class="card-header" style="background-color: #343b45;border-radius: 0px;">
                    <span style="color:white;font-weight: bold;font-size: 15px;">LOGIN AGEN PANEL</span>
                </div>
                <div class="card-body" style="padding:10px;">
                    <div class="mb-3">
                        <input on:change="{handleChange}"
                            bind:value={$form.username}
                            invalid={$errors.username.length > 0}
                            type="text"
                            class="form-control"
                            placeholder="Username"
                            aria-label="Username"/>
                    </div>
                    <div class="mb-3">
                        <input on:change="{handleChange}"
                            bind:value={$form.password}
                            invalid={$errors.password.length > 0}
                            type="password"
                            class="form-control"
                            placeholder="Password"
                            aria-label="Password"/>
                    </div>
                </div>
                <div class="card-footer" style="padding: 5px;margin:0px;border-radius: 0px;">
                    <div class="d-grid">
                        <button on:click={() => {
                                handleSubmit();
                            }} type="button" class="btn btn-dark">Submit</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>