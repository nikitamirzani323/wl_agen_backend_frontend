<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";

	export let table_header_font
	export let table_body_font
	export let token
	export let listAdmin = []
	export let listAdminrule = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "Admin"
    let sData = "";
    let username_flag = false;
    let myModal_newentry = ""
    let css_loader = "display: none;";
    let msgloader = "";
    let schema  = "";
    
    if(sData == "New"){
        schema = yup.object().shape({
            username: yup.string().required("The Username is required").
                        matches(/^[a-zA-z0-9]+$/, "Username must Character A-Z or a-z or 1-9 ").
                        min(3,"The Username minimal 3 Character").
                        max(30,"The Username mmaximal 30 Character"),
            password: yup.string().required("The Password is required").
                        min(4,"The Password minimal 3 Character").
                        max(30,"The Password mmaximal 30 Character"),
            rule: yup.string().required("The Rule is required"),
            name: yup.string().required("The Name is required").
                        matches(/^[a-zA-z0-9 ]+$/, "Name must Character A-Z or a-z or 1-9 ").
                        min(3,"The Name minimal 3 Character").
                        max(30,"The Name mmaximal 30 Character"),
            status: yup.string(),
        });
    }else if(sData == "Edit"){
        schema = yup.object().shape({
            username: yup.string(),
            password: yup.string().
                        max(30,"The Password mmaximal 30 Character"),
            rule: yup.string().required("The Rule is required"),
            name: yup.string().required("The Name is required").
                        matches(/^[a-zA-z0-9 ]+$/, "Name must Character A-Z or a-z or 1-9 ").
                        min(3,"The Name minimal 3 Character").
                        max(30,"The Name mmaximal 30 Character"),
            status: yup.string(),
        });
    }
    
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            username: "",
            password: "",
            rule:"",
            name:"",
            status:"Y"
        },
        validationSchema: schema,
        onSubmit:(values) => {
            handleSave(values.username,values.password,values.rule,values.name,values.status)
        }
    })
    const NewData = (e,username,pass,rule,name,status) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            username_flag = true
            $form.username = username
            $form.password = pass
            $form.rule = rule
            $form.name = name
            if(status == "ACTIVE"){
                status = "Y"
            }else{
                status = "N"
            }
            $form.status = status
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave(username,password,rule,name,status) {
        const res = await fetch("/api/saveadmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page:"ADMIN-SAVE",
                admin_rule: rule,
                admin_username: username,
                admin_password: password,
                admin_nama: name,
                admin_status: status,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            msgloader = json.message;
            myModal_newentry.hide()
            RefreshHalaman()
        } else if(json.status == 403){
            alert(json.message)
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    function clearField(){
        username_flag = false
        $form.username = ""
        $form.password = ""
        $form.rule = ""
        $form.name = ""
        $form.status = "Y"
    }
    $:{
        if(sData == "New"){
            if ($errors.username || $errors.password || $errors.rule || $errors.name || $errors.status){
                alert($errors.username+"\n"+$errors.password+"\n"+$errors.rule+"\n"+$errors.name+"\n"+$errors.status)
                $form.username = ""
                $form.password = ""
                $form.rule = ""
                $form.name = ""
                $form.status = "Y"
            }
        }else{
            if ($errors.username || $errors.rule || $errors.name || $errors.status){
                alert($errors.username+"\n"+$errors.rule+"\n"+$errors.name+"\n"+$errors.status)
                $form.username = ""
                $form.rule = ""
                $form.name = ""
                $form.status = "Y"
            }
        }
    }
    
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="NEW"
                button_title="New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-body">
                        <table class="table table-striped ">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                    <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">RULE</th>
                                    <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TIMEZONE</th>
                                    <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">JOIN DATE</th>
                                    <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">LAST LOGIN</th>
                                    <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">LAST IPADDRESS</th>
                                </tr>
                            </thead>
                            {#if totalrecord > 0}
                            <tbody>
                                {#each listAdmin as rec }
                                    <tr>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    NewData("Edit",rec.admin_username,"",rec.admin_rule,rec.admin_nama,rec.admin_status);
                                                }} 
                                                class="bi bi-pencil"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                            <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.admin_statuscss}">
                                                {rec.admin_status}
                                            </span>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_no}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_username}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_nama}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_rule}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_timezone}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_joindate}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_lastlogin}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_lastipaddres}</td>
                                    </tr>
                                {/each}
                            </tbody>
                            {:else}
                            <tbody>
                                <tr>
                                    <td colspan="10">
                                        <center>
                                            <Loader />
                                        </center>
                                    </td>
                                </tr>
                            </tbody>
                            {/if} 
                        </table>
                </slot:template>
            </Panel>

            
        </div>
    </div>
</div>

<Modal
	modal_id="modalentry"
	modal_size="modal-dialog-centered"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Rule</label>
			<select
                on:change="{handleChange}"
                bind:value={$form.rule} 
                class="form-control required">
                <option value="">--Select--</option>
                {#each listAdminrule as rec }
                    <option value="{rec.adminrule_idruleadmin}">{rec.adminrule_idruleadmin}</option>
                {/each}
            </select>
		</div>
		<div class="mb-3">
            <label for="exampleForm" class="form-label">Username</label>
            <input
                on:change="{handleChange}"
                bind:value={$form.username}
                invalid={$errors.username.length > 0}
                disabled='{username_flag}'
                type="text"
                class="form-control required"
                placeholder="Username"
                aria-label="Username"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Password</label>
            <input
                on:change="{handleChange}"
                bind:value={$form.password}
                invalid={$errors.password.length > 0}
                type="password"
                class="form-control required"
                placeholder="Password"
                aria-label="Password"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <input
                on:change="{handleChange}"
                bind:value={$form.name}
                invalid={$errors.name.length > 0}
                type="text"
                class="form-control required"
                placeholder="Name"
                aria-label="Name"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status</label>
            <select
                class="form-control required"
                on:change="{handleChange}"
                bind:value={$form.status}
                invalid={$errors.status.length > 0} >
                <option value="Y">ACTIVE</option>
                <option value="N">DEACTIVE</option>
            </select>
        </div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>