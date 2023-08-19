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
	let title_page = "TRANSAKSI"
    let listmember = [];
    let sData = "";
    let myModal_newentry = "";
    let flag_btnsave = true;
    let member_listbank = [];
    let member_id_field = "";
    let member_info_field = "";
    let bank_id_field = "";
    let bank_tipe_field = "";
    let bank_name_field = "";
    let bank_norek_field = "";
    let bank_status_field = "";
    let bank_create_field = "";
    let bank_update_field = "";
    let idrecord = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_idbanktype
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_nmrek
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    const ShowMember = () => {
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistmember"));
        myModal_newentry.show();
        member_listbank = [];
        member_id_field = ""
        member_info_field = ""
        call_member();
    };
    const InsertMember = (id,info,listbank) => {
        member_id_field = id
        member_info_field = info
        if (listbank != null) {
            for (var i = 0; i < listbank.length; i++) {
                member_listbank = [
                    ...member_listbank,
                    {
                        memberbank_id: listbank[i]["memberbank_id"],
                        memberbank_info: listbank[i]["memberbank_info"],
                    },
                ];
            }
        }
    };
    const NewDataDeposit = (e,id,idbank,tipe,norek,nmrek,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            idrecord = id
            bank_id_field = idbank;
            bank_tipe_field = tipe;
            bank_name_field = nmrek;
            bank_norek_field = norek;
            bank_status_field = status;
            bank_create_field = create;
            bank_update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycruddeposit"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    async function call_member() {
        listmember = [];
        const res = await fetch("/api/membersearch", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page:"PROVIDER-SAVE",
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listmember = [
                        ...listmember,
                        {
                            member_no: no,
                            member_id: record[i]["member_id"],
                            member_name: record[i]["member_name"],
                            member_listbank: record[i]["member_listbank"],
                        },
                    ];
                }
            }
        }
    }
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(bank_tipe_field == ""){
                flag = false
                msg += "The Tipe is required\n"
            }
            if(bank_norek_field == ""){
                flag = false
                msg += "The Nomor Rekening is required\n"
            }
            if(bank_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(bank_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(bank_tipe_field == ""){
                flag = false
                msg += "The Tipe is required\n"
            }
            if(bank_norek_field == ""){
                flag = false
                msg += "The Nomor Rekening is required\n"
            }
            if(bank_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(bank_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/banksave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"PROVIDER-SAVE",
                    agenbank_id: parseInt(idrecord),
                    agenbank_tipe: bank_tipe_field,
                    agenbank_idbanktype: bank_id_field,
                    agenbank_norek: bank_norek_field,
                    agenbank_nmrek: bank_name_field,
                    agenbank_status: bank_status_field,
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
    
   
    function clearField(){
        idrecord = "";
        bank_id_field = "";
        bank_tipe_field = "";
        bank_name_field = "";
        bank_norek_field = "";
        bank_create_field = "";
        bank_update_field = "";
    }
    function callFunction(event){
        switch(event.detail){
            case "NEWDEPOSIT":
                NewDataDeposit("New","","","");
                break;
            case "NEWWITHDRAW":
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
                button_function="NEWDEPOSIT"
                button_title="Deposit"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="NEWWITHDRAW"
                button_title="Withdraw"
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
                            placeholder="Search Bank"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TIPE</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">BANK</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                // e,id,idbank,tipe,norek,nmrek,status,create,update
                                                NewData("Edit",rec.home_id,  rec.home_idbanktype, rec.home_tipe, 
                                                rec.home_norek, rec.home_nmrek, rec.home_status, 
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_tipe}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        {rec.home_idbanktype} - {rec.home_norek} - {rec.home_nmrek} - {rec.home_status}
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
	modal_id="modalentrycruddeposit"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="Deposit / {sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Jenis Transaksi</label>
                    <input class="form-control"
                        type="text"
                        value="Deposit"
                        disabled
                        placeholder="Jenis Transaksi"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Member</label>
                    <div class="input-group mb-3">
                        <Input bind:value={member_info_field}
                            class="required"
                            type="text"
                            disabled
                            placeholder="Member"/>
                        <button on:click={() => {
                            ShowMember();
                            }} class="btn btn-primary" type="button" id="button-addon2">...</button>
                    </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Bank In</label>
                    <select
                        bind:value="{bank_id_field}" 
                        name="bankid" id="bankid" 
                        class="required form-control">
                        {#each member_listbank as rec}
                        <option value="{rec.memberbank_id}">{rec.memberbank_info}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Bank Out</label>
                    <div class="input-group mb-3">
                        <Input bind:value={bank_norek_field}
                            class="required"
                            type="text"
                            placeholder="Bank Out"/>
                        <button class="btn btn-primary" type="button" id="button-addon2">...</button>
                    </div>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Amount</label>
                    <Input bind:value={bank_name_field}
                        class="required"
                        style="text-align: right;"
                        type="text"
                        placeholder="Amount"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Note</label>
                    <textarea class="form-control" placeholder="Note" id="floatingTextarea"></textarea>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {bank_create_field}<br />
                        Update : {bank_update_field}
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
  modal_id="modallistmember"
  modal_size="modal-dialog-centered"
  modal_title="MEMBER"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_footer={false}>
  <slot:template slot="body">
    <table class="table table-sm">
      <thead>
        <tr>
            <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
            <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">MEMBER</th>
        </tr>
      </thead>
      <tbody>
        {#each listmember as rec}
          <tr on:click={() => {
                InsertMember(rec.member_id, rec.member_name, rec.member_listbank);
            }} style="cursor: pointer;">
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.member_no}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.member_name}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
</Modal>

