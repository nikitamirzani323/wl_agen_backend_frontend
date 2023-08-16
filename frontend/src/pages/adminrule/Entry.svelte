<script>
    import Loader from "../../components/Loader.svelte";
    import Button from "../../components/Button.svelte";
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    export let sData = "";
    export let token = "";
    export let adminrule_idadmin = 0;
    export let adminrule_name = "";
    export let adminrule_rule = "";
    let adminrule_rule_field = adminrule_rule;

    let css_loader = "display: none;";
    let msgloader = "";
    let flag_name = true
    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        admin_name_field: yup
            .string()
            .required("Name is Required")
            .matches(
                /^[a-zA-z0-9 ]+$/,
                "Name must Character A-Z or a-z or 1-9 or space"
            )
            .min(4, "Name must be at least 4 Character")
            .max(70, "Name must be at most 4 Character"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            admin_name_field: adminrule_name,
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(values.admin_name_field);
        },
    });
    $: {
        if ($errors.admin_name_field) {
            alert($errors.admin_name_field);
            form.admin_name_field = "";
        }
    }
    const BackHalaman = () => {
        dispatch("handleBackHalaman", "call");
    };
    async function SaveTransaksi(name) {
        let temp_name = name.toLowerCase()
        let temp_name2 = temp_name.replace(/\s/g, '')
        const res = await fetch("/api/saveadminrule", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "ADMINRULE-SAVE",
                agenadminrule_id: parseInt(adminrule_idadmin),
                agenadminrule_name: temp_name2,
                agenadminrule_rule: adminrule_rule_field.toString(),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            msgloader = json.message;
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    async function Updateconfig() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (adminrule_rule_field == "") {
            flag = true;
            msg += "The List is required\n";
        }
        if (flag == false) {
            let temp_name = adminrule_name.toLowerCase()
            let temp_name2 = temp_name.replace(/\s/g, '')
            const res = await fetch("/api/saveadminrule", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "ADMINRULE-SAVE",
                    agenadminrule_id: parseInt(adminrule_idadmin),
                    agenadminrule_name: temp_name2,
                    agenadminrule_rule: adminrule_rule_field.toString(),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
    function callFunction(event) {
        switch (event.detail) {
            case "BACK":
                BackHalaman();
                break;
            case "REFRESH":
                RefreshHalaman();
                break;
            case "SAVE":
                handleSubmit();
                break;
        }
    }
    function lowerCase(element) {
        function onInput(event) {
            element.value = element.value.toLowerCase();
            element.value = element.value.replace(/\s/g, '');
        }
        element.addEventListener("input", onInput);
        return {
            destroy() {
                element.removeEventListener("input", onInput);
            },
        };
    }
    if(adminrule_name == "master"){
        flag_name = false;
    }
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-6">
            <Button
                on:click={callFunction}
                button_function="BACK"
                button_title="Back"
                button_css="btn-primary"/>
        </div>
    </div>
    <div class="clearfix" />
    <div class="row">
        <div class="col-sm-3">
            <div class="card" style="border-radius: 0px;margin-top:10px;">
                <div class="card-header" style="">
                    Admin Rule / {sData}
                    <div class="float-end">
                        <button
                            on:click={() => {
                                handleSubmit();
                            }}
                            class="btn btn-warning btn-sm"
                            style="border-radius: 0px;">
                            Save
                        </button>
                    </div>
                </div>
                <div class="card-body" style="height:400px;">
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">Name</label>
                        {#if flag_name}
                        <input
                            use:lowerCase 
                            on:change={handleChange}
                            bind:value={$form.admin_name_field}
                            invalid={$errors.admin_name_field.length > 0}
                            type="text"
                            maxlength="70"
                            class="form-control"
                            placeholder="Name"
                            aria-label="Name"/>
                        {:else}
                        <input
                            disabled
                            on:change={handleChange}
                            bind:value={$form.admin_name_field}
                            invalid={$errors.admin_name_field.length > 0}
                            type="text"
                            maxlength="70"
                            class="form-control"
                            placeholder="Name"
                            aria-label="Name"/>
                        {/if}
                    </div>
                </div>
            </div>
        </div>
        <div class="col-sm-9">
            <div class="card" style="border-radius: 0px;margin-top:10px;">
                <div class="card-header" style="">
                    Setting RULE
                    <div class="float-end">
                        <button
                            on:click={() => {
                                Updateconfig();
                            }}
                            class="btn btn-warning btn-sm"
                            style="border-radius: 0px;">
                            Save
                        </button>
                    </div>
                </div>
                <div class="card-body" style="height:700px;overflow-y: scroll;">
                    <table class="table">
                        <thead>
                            <tr>
                                <th colspan="2">DASHBOARD</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="DASHBOARD-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                            </tr>
                        </tbody>
                    </table>
                  
                    <table class="table">
                        <thead>
                            <tr>
                                <th colspan="2">MEMBER MANAGEMENT</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td width="1%">
                                    <input bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="MEMBER-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                            <tr>
                                <td width="1%">
                                    <input bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="MEMBER-SAVE"/>
                                </td>
                                <td width="*">SAVE</td>
                            </tr>
                        </tbody>
                    </table>
                    <table class="table">
                        <thead>
                            <tr>
                                <th colspan="2">ADMIN MANAGEMENT</th>
                                <th colspan="2">ADMIN RULE</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td width="1%">
                                    <input bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMIN-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                                <td width="1%">
                                    <input bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMINRULE-VIEW"/>
                                </td>
                                <td width="*">VIEW</td>
                            </tr>
                            <tr>
                                <td width="1%">
                                    <input bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMIN-SAVE"/>
                                </td>
                                <td width="*">SAVE</td>
                                <td width="1%">
                                    <input bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMINRULE-SAVE"/>
                                </td>
                                <td width="*">SAVE</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</div>
