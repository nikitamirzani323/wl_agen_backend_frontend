<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

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
    let idrecord = "";
    let idrule_field = 0;
    let username_field = "";
    let password_field = "";
    let name_field = "";
    let phone1_field = "";
    let phone2_field = "";
    let status_field = "";
    let myModal_newentry = ""
    let css_loader = "display: none;";
    let msgloader = "";
    
    const NewData = (e,id,idrule,username,pass,name,phone1,phone2,status) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            username_flag = true
            idrecord = id
            idrule_field = idrule
            username_field = username;
            password_field = pass;
            name_field = name;
            phone1_field = phone1;
            phone2_field = phone2;
            status_field = status;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(idrule_field == ""){
                flag = false
                msg += "The ID Rule is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone 1 is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idrule_field == ""){
                flag = false
                msg += "The ID Rule is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone 1 is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/saveadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"ADMIN-SAVE",
                    admin_id: idrecord,
                    admin_idrule: parseInt(idrule_field),
                    admin_username: username_field,
                    admin_password: password_field,
                    admin_nama: name_field,
                    admin_phone1: phone1_field,
                    admin_phone2: phone2_field,
                    admin_status: status_field,
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
        }else{
            alert(msg)
        }
    }
    function clearField(){
        idrecord = ""
        idrule_field = 0
        username_flag = false
        username_field = "";
        password_field = "";
        name_field = "";
        phone1_field = "";
        phone2_field = "";
        status_field = "";
    }
    $:{
        
    }
    
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSave();break;
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
                                    <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TIPE</th>
                                    <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">RULE</th>
                                    <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                                    <th NOWRAP width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
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
                                                    //e,id,idrule,username,pass,name,phone1,phone2,status
                                                    NewData("Edit",rec.home_id,rec.home_idrule,rec.home_username,"",
                                                    rec.home_nama,rec.home_phone1,rec.home_phone2,
                                                    rec.home_status);
                                                }} 
                                                class="bi bi-pencil"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                            <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_statuscss}">
                                                {rec.home_status}
                                            </span>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_tipe}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmrule}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_username}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nama}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_phone1}/{rec.home_phone2}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_lastlogin}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_lastipaddres}</td>
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
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Rule</label>
                    <select
                        bind:value={idrule_field} 
                        class="form-control required">
                        <option value="">--Select--</option>
                        {#each listAdminrule as rec }
                            <option value="{rec.adminrule_idruleadmin}">{rec.adminrule_nmruleadmin}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Username</label>
                    <input
                        bind:value={username_field}
                        disabled='{username_flag}'
                        type="text"
                        class="form-control required"
                        placeholder="Username"
                        aria-label="Username"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Password</label>
                    <input
                        bind:value={password_field}
                        type="password"
                        class="form-control required"
                        placeholder="Password"
                        aria-label="Password"/>
                </div>
            </div>
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <input
                        bind:value={name_field}
                        type="text"
                        class="form-control required"
                        placeholder="Name"
                        aria-label="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 1</label>
                    <input
                        bind:value={phone1_field}
                        type="text"
                        class="form-control required"
                        placeholder="Phone 1"
                        aria-label="Phone1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 2</label>
                    <input
                        bind:value={phone2_field}
                        type="text"
                        class="form-control required"
                        placeholder="Phone 2"
                        aria-label="Phone2"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select class="form-control required"
                        bind:value={status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
            </div>
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