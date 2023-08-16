<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "CATEGORY BANK"
	let title_detail_page = ""
    let sData = "";
    let sDataDetail = "";
    let myModal_newentry = "";
    let name_field = "";
    let status_field = "N";
    let create_field = "";
    let update_field = "";
    let flag_btnsave = true;
    let flag_idrecorddetail_field = false;
    let flag_detail_btnsave = true;
    let name_detail_field = "";
    let img_detail_field = "";
    let status_detail_field = "N";
    let create_detail_field = "";
    let update_detail_field = "";
    let idrecord = "";
    let idrecorddetail = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id,nmcatebank,statuscatebank,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            idrecord = parseInt(id)
            name_field = nmcatebank;
            status_field = statuscatebank;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const NewDataDetail = (e,id,name,iddetail,nmdetail,imgdetail,statusdetail,createdetail,updatedetail) => {
        sDataDetail = e
        title_detail_page = name
        idrecord = id
        if(sDataDetail == "New"){
            cleardetailField()
        }else{
            flag_idrecorddetail_field = true;
            idrecorddetail = iddetail;
            name_detail_field = nmdetail;
            img_detail_field = imgdetail;
            status_detail_field = statusdetail;
            create_detail_field = createdetail;
            update_detail_field = updatedetail;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrydetailcrud"));
        myModal_newentry.show();
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idrecord.length == 4){
                flag = false
                msg += "The ID is maxlengt 4\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Domain is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/catebanksave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CATEBANK-SAVE",
                    catebank_id: parseInt(idrecord),
                    catebank_name: name_field,
                    catebank_status: status_field,
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
    async function handleDetailSave() {
        let flag = true
        let msg = ""
        if(sDataDetail == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(name_detail_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(status_detail_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(name_detail_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(status_detail_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_detail_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/banktypesave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataDetail,
                    page:"CATEBANK-SAVE",
                    banktype_id: idrecorddetail.toUpperCase(),
                    banktype_idcatebank: parseInt(idrecord),
                    banktype_name: name_detail_field.toUpperCase(),
                    banktype_img: img_detail_field,
                    banktype_status: status_detail_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                if(sDataDetail=="New"){
                    cleardetailField()
                }
                flag_detail_btnsave = true;
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                alert(json.message)
                flag_detail_btnsave = true;
            } else {
                msgloader = json.message;
                flag_detail_btnsave = true;
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
        name_field = "";
        status_field = "N";
        create_field = "";
        update_field = "";
        name_detail_field = "";
        img_detail_field = "";
        status_detail_field = "N";
        create_detail_field = "";
        update_detail_field = "";
    }
    function cleardetailField(){
        flag_idrecorddetail_field = false;
        idrecorddetail = "";
        name_detail_field = "";
        img_detail_field = "";
        status_detail_field = "N";
        create_detail_field = "";
        update_detail_field = "";
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
            case "SAVEDETAIL":
                handleDetailSave();break;
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
    function uperCase(element) {
        function onInput(event) {
            element.value = element.value.toUpperCase();
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
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Category Bank"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="2">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CATEGORY BANK</th>
                                <th NOWRAP width="35%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">BANK + PROVIDER</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                NewData("Edit",rec.home_id, rec.home_name, rec.home_status, rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                NewDataDetail("New",rec.home_id, rec.home_name);
                                            }} class="bi bi-plus-circle"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {rec.home_status}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {#each rec.home_list as rec2}
                                            <i on:click={() => {
                                                NewDataDetail("Edit",rec.home_id, rec.home_name, 
                                                rec2.banktype_id, rec2.banktype_name, rec2.banktype_img,rec2.banktype_status,
                                                rec2.banktype_create, rec2.banktype_update);
                                            }} class="bi bi-pencil" style="cursor: pointer;"></i>
                                            {rec2.banktype_status}
                                            <a href="{rec2.banktype_img}" target="_blank">{rec2.banktype_name}</a><br />
                                        {/each}
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
	modal_size="modal-dialog-centered"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Category Bank</label>
            <input bind:value={name_field}
                use:uperCase
                class="required form-control"
                maxlength="70" 
                type="text"
                placeholder="Name"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status</label>
            <select class="form-control required" bind:value="{status_field}">
                <option value="Y">Y</option>
                <option value="N">N</option>
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
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave == true}
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
	modal_id="modalentrydetailcrud"
	modal_size="modal-dialog-centered"
	modal_title="List Bank + Provider {title_detail_page+"/"+sDataDetail}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">ID</label>
            {#if flag_idrecorddetail_field == true}
            <input bind:value={idrecorddetail}
                use:uperCase  
                disabled 
                class="required form-control"
                maxlength="10"
                type="text"
                placeholder="ID"/>
            {:else}
            <input bind:value={idrecorddetail}
                use:uperCase  
                class="required form-control"
                maxlength="10"
                type="text"
                placeholder="ID"/>
            {/if}
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <input bind:value={name_detail_field}
                use:uperCase
                class="required form-control"
                maxlength="70" 
                type="text"
                placeholder="Name"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Image</label>
            <input bind:value={img_detail_field}
                class="required form-control"
                maxlength="350" 
                type="text"
                placeholder="Image"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status</label>
            <select class="form-control required" bind:value="{status_detail_field}">
                <option value="Y">Y</option>
                <option value="N">N</option>
            </select>
        </div>
        {#if sDataDetail != "New"}
        <div class="mb-3">
            <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                Create : {create_detail_field}<br />
                Update : {update_detail_field}
            </div>
        </div>
        {/if}
	</slot:template>
	<slot:template slot="footer">
        {#if flag_detail_btnsave == true}
        <Button on:click={() => {
                handleDetailSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

