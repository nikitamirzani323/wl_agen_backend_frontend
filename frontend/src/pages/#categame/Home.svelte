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
	export let listProvider = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "CATEGORY GAME";
    let title_detail_page = "";
    let sData = "";
    let sDataDetail = "";
    let myModal_newentry = "";
    let flag_id_field = false;
    let flag_btnsave = true;
    let flag_idrecorddetail_field = false;
    let name_field = "";
    let status_field = "";
    let create_field = "";
    let update_field = "";

    let idprovider_detail_field = 0;
    let idcategorygame_detail_field = "";
    let name_detail_field = "";
    let image_detail_field = "";
    let multiplier_detail_field = 1;
    let urlstaging_detail_field = "";
    let urlproduction_detail_field = "";
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
    
    const NewData = (e,id,name,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            flag_id_field = true;
            idrecord = id
            name_field = name;
            status_field = status;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const NewDataDetail = (e,id,idcategame,idprovider,name,img,multiplier,urlstaging,urlproduction,status,create,update) => {
        sDataDetail = e
        title_detail_page = name
        idrecord = id
        if(sDataDetail == "New"){
            cleardetailField()
            idcategorygame_detail_field = idcategame;
        }else{
            flag_idrecorddetail_field = true;
            idrecorddetail = id;
            idprovider_detail_field = idprovider;
            idcategorygame_detail_field = idcategame;
            name_detail_field = name;
            image_detail_field = img;
            multiplier_detail_field = multiplier;
            urlstaging_detail_field = urlstaging;
            urlproduction_detail_field = urlproduction;
            status_detail_field = status;
            create_detail_field = create;
            update_detail_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrygamecrud"));
        myModal_newentry.show();
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
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
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
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
            let temp_id = idrecord.replace(/\s/g, '');
            const res = await fetch("/api/categamesave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"PROVIDER-SAVE",
                    categame_id: temp_id.toUpperCase(),
                    categame_name: name_field,
                    categame_status: status_field,
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
    async function handledetailSave() {
        let flag = true
        let msg = ""
        if(sDataDetail == "New"){
            if(idprovider_detail_field == ""){
                flag = false
                msg += "The Provider is required\n"
            }
            if(idprovider_detail_field == 0){
                flag = false
                msg += "The Provider is required\n"
            }
            if(name_detail_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(multiplier_detail_field < 1){
                flag = false
                msg += "The Multiplier cannot 0, minimal 1\n"
            }
            if(image_detail_field == ""){
                flag = false
                msg += "The Image is required\n"
            }
            if(urlstaging_detail_field == ""){
                flag = false
                msg += "The URL Staging is required\n"
            }
            if(urlproduction_detail_field == ""){
                flag = false
                msg += "The URL Production is required\n"
            }
            if(status_detail_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecorddetail == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idprovider_detail_field == ""){
                flag = false
                msg += "The Provider is required\n"
            }
            if(idprovider_detail_field == 0){
                flag = false
                msg += "The Provider is required\n"
            }
            if(name_detail_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(multiplier_detail_field < 1){
                flag = false
                msg += "The Multiplier cannot 0, minimal 1\n"
            }
            if(image_detail_field == ""){
                flag = false
                msg += "The Image is required\n"
            }
            if(urlstaging_detail_field == ""){
                flag = false
                msg += "The URL Staging is required\n"
            }
            if(urlproduction_detail_field == ""){
                flag = false
                msg += "The URL Production is required\n"
            }
            if(status_detail_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/gamesave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataDetail,
                    page:"CATEGAME-SAVE",
                    game_id: idrecorddetail,
                    game_idcategame: idcategorygame_detail_field,
                    game_idprovider: parseInt(idprovider_detail_field),
                    game_name: name_detail_field,
                    game_img: image_detail_field,
                    game_multiplier: parseInt(multiplier_detail_field),
                    game_urlstaging: urlstaging_detail_field,
                    game_urlproduction: urlproduction_detail_field,
                    game_status: status_detail_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataDetail=="New"){
                    cleardetailField()
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
        flag_id_field = false
        name_field = "";
        status_field = "";
        create_field = "";
        update_field = "";
    }
    function cleardetailField(){
        flag_idrecorddetail_field = false;
        idrecorddetail = "";
        idprovider_detail_field = "";
        idcategorygame_detail_field = "";
        name_detail_field = "";
        image_detail_field = "";
        multiplier_detail_field = 1;
        urlstaging_detail_field = "";
        urlproduction_detail_field = "";
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
    const handleKeyboard_format = () => {
        let numbera;

        for (let i = 0; i < multiplier_detail_field.length; i++) {
            numbera = parseInt(multiplier_detail_field[i]);
            if (isNaN(numbera)) {
                if (multiplier_detail_field[i] != ":") {
                    multiplier_detail_field = "";
                }
            }
        }
    };
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
    }
    function upperCase(element) {
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
                            placeholder="Search Category Game"
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
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CATEGORY GAME</th>
                                <th NOWRAP width="45%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">GAME</th>
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
                                           
                                                NewDataDetail("New","",rec.home_id);
                                            }} class="bi bi-plus-circle"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {#if rec.home_list != null}
                                        <table class="table table-striped-columns">
                                            <thead>
                                                <tr>
                                                    <th style="text-align: center;vertical-align: top;font-size: 11px;">STATUS</th>
                                                    <th style="text-align: left;vertical-align: top;font-size: 11px;">GAME</th>
                                                    <th style="text-align: left;vertical-align: top;font-size: 11px;">PROVIDER</th>
                                                    <th NOWRAP style="text-align: left;vertical-align: top;font-size: 11px;">URL STAGING</th>
                                                    <th NOWRAP style="text-align: left;vertical-align: top;font-size: 11px;">URL PRODUCTION</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                {#each rec.home_list as rec2}
                                                <tr>
                                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec2.game_status_css}">
                                                            {status(rec2.game_status)}
                                                        </span>
                                                    </td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">
                                                        <span
                                                            on:click={() => {
                                                                 // e,id,idcategame,idprovider,name,img,multiplier,urlstaging,urlproduction,status,create,update
                                                                NewDataDetail("Edit",rec2.game_id,rec2.game_idcategame,
                                                                rec2.game_idprovider,rec2.game_name,
                                                                rec2.game_img,rec2.game_multiplier,
                                                                rec2.game_urlstaging,rec2.game_urlproduction,rec2.game_status,
                                                                rec2.game_create,rec2.game_update);
                                                            }} 
                                                            style="cursor: pointer;text-decoration:underline;">{rec2.game_name}</span>
                                                    </td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">{rec2.game_nmprovider}</td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">{rec2.game_urlstaging}</td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">{rec2.game_urlproduction}</td>
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
	modal_size="modal-dialog-centered "
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">CODE</label>
            {#if flag_id_field == true}
            <input bind:value={idrecord}
                use:upperCase  
                disabled
                class="required form-control"
                maxlength="20"
                type="text"
                placeholder="CODE"/>
            {:else}
            <input bind:value={idrecord}
                use:upperCase
                class="required form-control"
                maxlength="20"
                type="text"
                placeholder="CODE"/>
            {/if}
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <Input bind:value={name_field}
                class="required"
                type="text"
                placeholder="Name"/>
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
	modal_id="modalentrygamecrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{idcategorygame_detail_field+"/"+sDataDetail}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Category Game</label>
                    <Input bind:value={idcategorygame_detail_field}
                        disabled
                        class="required"
                        type="text"
                        placeholder="Category Game"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Provider</label>
                    <select
                        bind:value="{idprovider_detail_field}" 
                        name="provider" id="provider" 
                        class="required form-control">
                        {#each listProvider as rec}
                        <option value="{rec.provider_id}">{rec.provider_name}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input bind:value={name_detail_field}
                        class="required"
                        type="text"
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Multiplier</label>
                    <Input
                      bind:value={multiplier_detail_field}
                      on:keyup={handleKeyboard_format}
                      class="required"
                      type="text"
                      style="text-align:right;"
                      placeholder="Display"/>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Image</label>
                    <Input bind:value={image_detail_field}
                        class="required"
                        type="text"
                        placeholder="Image"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">URL Staging</label>
                    <Input bind:value={urlstaging_detail_field}
                        class="required"
                        type="text"
                        placeholder="URL Staging"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">URL Production</label>
                    <Input bind:value={urlproduction_detail_field}
                        class="required"
                        type="text"
                        placeholder="URL Production"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_detail_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
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
            </div>
        </div>
        
        
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handledetailSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>
