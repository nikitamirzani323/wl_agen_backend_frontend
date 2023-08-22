<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";
  import { intros } from "svelte/internal";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let listBank = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "MEMBER"
    let sData = "";
    let sDataBank = "";
    let myModal_newentry = "";
    let flag_btnsave = true;
    let username_flag = false;
    let username_field = "";
    let password_field = "";
    let name_field = "";
    let email_field = "";
    let phone_field = "";
    let status_field = "";
    let create_field = "";
    let update_field = "";
    let bank_idmember_field = "";
    let bank_id_field = "";
    let bank_name_field = "";
    let bank_norek_field = "";
    let idrecord = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_username
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id,username,name,email,phone,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            username_flag = true;
            idrecord = id
            username_field = username;
            name_field = name;
            email_field = email;
            phone_field = phone;
            status_field = status;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const NewDataBank = (e,idmember,idbanktype,norek,name) => {
        sDataBank = e
        bank_idmember_field = idmember;
        if(sDataBank == "New"){
            clearFieldBank()
        }else{
            bank_idmember_field = idmember;
            bank_id_field = idbanktype;
            bank_name_field = name;
            bank_norek_field = norek;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrudbank"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(username_field == ""){
                flag = false
                msg += "The Username is required\n"
            }
            if(username_field.length < 4){
                flag = false
                msg += "The Username is minglength 4\n"
            }
            if(password_field == ""){
                flag = false
                msg += "The Password is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(phone_field == ""){
                flag = false
                msg += "The Phone is required\n"
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
            if(username_field == ""){
                flag = false
                msg += "The Username is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(phone_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/membersave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"PROVIDER-SAVE",
                    member_id: idrecord,
                    member_username: username_field,
                    member_password: password_field,
                    member_name: name_field,
                    member_phone: phone_field,
                    member_email: email_field,
                    member_status: status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleBankSave() {
        let flag = true
        let msg = ""
        if(sDataBank == "New"){
            if(bank_idmember_field == ""){
                flag = false
                msg += "The Code Member is required\n"
            }
            if(bank_id_field == ""){
                flag = false
                msg += "The Bank is required\n"
            }
            if(bank_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(bank_norek_field == ""){
                flag = false
                msg += "The Norek is required\n"
            }
        }else{
            if(bank_idmember_field == ""){
                flag = false
                msg += "The Code Member is required\n"
            }
            if(bank_id_field == ""){
                flag = false
                msg += "The Bank is required\n"
            }
            if(bank_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(bank_norek_field == ""){
                flag = false
                msg += "The Norek is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/memberbanksave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataBank,
                    page:"PROVIDER-SAVE",
                    memberbank_idagenmember: bank_idmember_field,
                    memberbank_idbanktype: bank_id_field,
                    memberbank_norek: bank_norek_field,
                    memberbank_nmownerbank: bank_name_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataBank=="New"){
                    clearFieldBank()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleBankDelete(idmemberbank,idagenmember) {
        let flag = true
        let msg = ""
        if(idagenmember == ""){
            flag = false
            msg += "The Code Member is required\n"
        }
        if(idmemberbank == ""){
            flag = false
            msg += "The Bank is required\n"
        }
       
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/memberbankdelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: "New",
                    page:"PROVIDER-SAVE",
                    memberbank_id: parseInt(idmemberbank),
                    memberbank_idagenmember: idagenmember,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataBank=="New"){
                    clearFieldBank()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
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
        idrecord = "";
        username_flag = false;
        username_field = "";
        password_field = "";
        name_field = "";
        email_field = "";
        phone_field = "";
        status_field = "";
        create_field = "";
        update_field = "";
        bank_id_field = "";
        bank_name_field = "";
        bank_norek_field = "";
    }
    function clearFieldBank(){
        bank_id_field = "";
        bank_name_field = "";
        bank_norek_field = "";
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterTafsirMimpi = [];
                listHome = [];
                const tafsir = {
                    searchTafsirMimpi,
                };
                dispatch("handleTafsirMimpi", tafsir);
        }  
    };
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
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
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Member"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan=2>&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CODE</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CREDIT</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">INFO</th>
                                <th NOWRAP width="25%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">BANK</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                // e,id,username,name,email,phone,status,create,update
                                                NewData("Edit",rec.home_id, rec.home_username, 
                                                rec.home_name, rec.home_email, rec.home_phone, rec.home_status, 
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                // e,id,idmember,idbanktype,norek,name
                                                NewDataBank("New",rec.home_id);
                                            }} class="bi bi-database-add"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.home_credit_css}">{new Intl.NumberFormat().format(rec.home_credit)}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_username}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        NAME : {rec.home_name}<br />
                                        EMAIL : {rec.home_email}<br />
                                        PHONE : {rec.home_phone}<br />
                                        LASTLOGIN : {rec.home_lastlogin}<br />
                                        LASTIPADDRESS : {rec.home_ipaddress}<br />
                                        LASTTIMEZONE : {rec.home_timezone}
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {#if rec.home_listbank != null}
                                        <table class="table table-bordered">
                                            <tbody>
                                                {#each rec.home_listbank as rec2}
                                                <tr>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;cursor:pointer;">
                                                        <i on:click={() => {
                                                            //e,id,idmaster,tipe,username,name,phone1,phone2,status,create,update
                                                            handleBankDelete(rec2.memberbank_id,rec.home_id);
                                                        }} class="bi bi-trash"></i>
                                                    </td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">{rec2.memberbank_idbanktype}</td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">{rec2.memberbank_norek}</td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">{rec2.memberbank_nmownerbank}</td>
                                                </tr>
                                                {/each}
                                            </tbody>
                                        </table>
                                        {/if}
                                    </td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
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
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Username</label>
                    <input 
                        use:lowerCase
                        bind:value={username_field} 
                        disabled='{username_flag}'
                        class="form-control required"
                        minlength="4"
                        maxlength="20"
                        type="text"
                        placeholder="Username"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Password</label>
                    <input
                        bind:value={password_field} 
                        type="password"
                        maxlength="30"
                        class="form-control "
                        placeholder="Password"
                        aria-label="Password"/>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input bind:value={name_field}
                        class="required"
                        type="text"
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Email</label>
                    <Input bind:value={email_field}
                        class=""
                        type="text"
                        placeholder="Email"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone</label>
                    <Input bind:value={phone_field}
                        class="required"
                        type="text"
                        placeholder="Phone"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {create_field}<br />
                        Update : {update_field}
                    </div>
                </div>
                {/if}
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalentrycrudbank"
	modal_size="modal-dialog-centered"
	modal_title="Bank {bank_idmember_field+"/"+sDataBank}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Bank</label>
            <select
                bind:value="{bank_id_field}" 
                name="bankid" id="bankid" 
                class="required form-control">
                {#each listBank as rec}
                <option value="{rec.bank_id}">{rec.bank_category} - {rec.bank_id}</option>
                {/each}
            </select>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Nomor Rekening</label>
            <Input bind:value={bank_norek_field}
                class="required"
                type="text"
                placeholder="Bank Nomor Rekening"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Nama Bank</label>
            <Input bind:value={bank_name_field}
                class="required"
                type="text"
                placeholder="Bank Nama Pemilik Rekening"/>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleBankSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

