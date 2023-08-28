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
    let sData = "";
    let myModal_newentry = "";
    let flag_btnsave = true;

    let filter_listmember = [];
    let listmember = [];
    let member_listbank = [];
    let member_search = "";
    let member_id_field = "";
    let member_info_field = "";

    let filter_listagenbank = [];
    let listagenbank = [];
    let agenbank_search = "";
    let agenbank_id_field = "";
    let agenbank_info_field = "";
    
    let temp_credit_member = 0;
    let transaksi_temp_field = "";
    let transaksi_id_field = "";
    let transaksi_idmember_field = "";
    let transaksi_nmmember_field = "";
    let transaksi_tipe_field = "";
    let transaksi_bankin_id_field = "";
    let transaksi_bankin_info_field = "";
    let transaksi_bankout_id_field = "";
    let transaksi_bankout_info_field = "";
    let transaksi_amount_field = 0;
    let transaksi_note_field = "";
    let transaksi_status_field = "";
    let transaksi_create_field = "";
    let transaksi_update_field = "";
    let idrecord = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_nmmember
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
        if (member_search) {
            filter_listmember = listmember.filter(
                (item) =>
                    item.member_name
                        .toLowerCase()
                        .includes(member_search.toLowerCase()) 
            );
        } else {
            filter_listmember = [...listmember];
        }
        if (agenbank_search) {
            filter_listagenbank = listagenbank.filter(
                (item) =>
                    item.agenbank_info
                        .toLowerCase()
                        .includes(agenbank_search.toLowerCase()) 
            );
        } else {
            filter_listagenbank = [...listagenbank];
        }
    }
    const ShowInOut = (e) => {
        switch(transaksi_tipe_field){
            case "DEPOSIT":
                transaksi_temp_field = e;
                if(e == "IN"){//FORM IN
                    myModal_newentry = new bootstrap.Modal(document.getElementById("modallistagenbank"));
                    myModal_newentry.show();
                    filter_listagenbank = [];
                    listagenbank = [];

                    transaksi_bankin_id_field = "";
                    transaksi_bankin_info_field = "";
                    call_agenbank();
                }else{//FORM OUT
                    myModal_newentry = new bootstrap.Modal(document.getElementById("modallistmember"));
                    myModal_newentry.show();
                    filter_listmember = [];

                    transaksi_bankout_id_field = "";
                    transaksi_bankout_info_field = "";
                    call_member();
                }
                break;
            case "WITHDRAW":
                transaksi_temp_field = e;
                if(e == "OUT"){//FORM OUT
                    myModal_newentry = new bootstrap.Modal(document.getElementById("modallistagenbank"));
                    myModal_newentry.show();
                    filter_listagenbank = [];
                    listagenbank = [];

                    transaksi_bankout_id_field = "";
                    transaksi_bankout_info_field = "";
                    call_agenbank();
                }else{//FORM IN
                    myModal_newentry = new bootstrap.Modal(document.getElementById("modallistmember"));
                    myModal_newentry.show();
                    filter_listmember = [];
                    member_listbank = [];
                    transaksi_bankin_id_field = "";
                    transaksi_bankin_info_field = "";
                    call_member();
                }
                break;
        }
        
    };
    const InsertInOut = (id,info,listbank,credit) => {
        switch(transaksi_tipe_field){
            case "DEPOSIT":
                if(transaksi_temp_field == "IN"){
                    transaksi_bankin_id_field = id
                    transaksi_bankin_info_field = info
                }else{
                    transaksi_idmember_field = id
                    transaksi_bankout_info_field = info
                    member_listbank = []
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
                }
                break;
            case "WITHDRAW":
                temp_credit_member = credit
                if(transaksi_temp_field == "OUT"){
                    transaksi_bankout_id_field = id
                    transaksi_bankout_info_field = info
                }else{
                    transaksi_idmember_field = id
                    transaksi_bankin_info_field = info
                    member_listbank = []
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
                }
                break;
        }
        myModal_newentry.hide();
    };
    const NewDataTransaksi = (e,tipe,id,idmember,nmmember,bank_in,bank_out,bank_in_info,bank_out_info,amount,status,note,create,update) => {
        sData = e
        if(e=="New"){
            clearField()
            transaksi_tipe_field = tipe
            transaksi_status_field = "PROCESS"
            myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycruddeposit"));
            myModal_newentry.show();
        }else{
            transaksi_id_field = id;
            transaksi_idmember_field = idmember;
            transaksi_nmmember_field = nmmember;
            transaksi_tipe_field = tipe;
            transaksi_bankin_id_field = bank_in;
            transaksi_bankin_info_field = bank_in_info;
            transaksi_bankout_id_field = bank_out;
            transaksi_bankout_info_field = bank_out_info;
            transaksi_amount_field = amount;
            transaksi_note_field = note;
            transaksi_status_field = status;
            transaksi_create_field = create;
            transaksi_update_field = update;
            if(tipe == "DEPOSIT"){
                myModal_newentry = new bootstrap.Modal(document.getElementById("modalformdeposit"));
                myModal_newentry.show();
            }else{
                myModal_newentry = new bootstrap.Modal(document.getElementById("modalformwithdraw"));
                myModal_newentry.show();
            }
            
        }
        
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function call_agenbank() {
        listagenbank = [];
        const res = await fetch("/api/banklist", {
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
                    listagenbank = [
                        ...listagenbank,
                        {
                            agenbank_no: no,
                            agenbank_id: record[i]["agenbank_id"],
                            agenbank_info: record[i]["agenbank_info"],
                        },
                    ];
                }
            }
        }
    }
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
                            member_credit: record[i]["member_credit"],
                            member_listbank: record[i]["member_listbank"],
                        },
                    ];
                }
            }
        }
    }
   
    async function handleSaveTransaksi() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(transaksi_tipe_field == ""){
                flag = false
                msg += "The Jenis Transaksi is required\n"
            }
            if(transaksi_idmember_field == ""){
                flag = false
                msg += "The Member is required\n"
            }
            if(transaksi_bankin_id_field == ""){
                flag = false
                msg += "The Bank In is required\n"
            }
            if(transaksi_bankout_id_field == ""){
                flag = false
                msg += "The Bank out is required\n"
            }
            if(transaksi_amount_field == ""){
                flag = false
                msg += "The Amount is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(transaksi_tipe_field == ""){
                flag = false
                msg += "The Jenis Transaksi is required\n"
            }
            if(transaksi_idmember_field == ""){
                flag = false
                msg += "The Member is required\n"
            }
            if(transaksi_bankin_id_field == ""){
                flag = false
                msg += "The Bank In is required\n"
            }
            if(transaksi_bankout_id_field == ""){
                flag = false
                msg += "The Bank out is required\n"
            }
            if(transaksi_amount_field == ""){
                flag = false
                msg += "The Amount is required\n"
            }
        }
        if(transaksi_tipe_field == "WITHDRAW"){
            if(parseInt(temp_credit_member) < parseInt(transaksi_amount_field)){
                flag = false
                msg += "The Amount exceed Credit Member\n"
            }
        }
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/transdpwdsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"PROVIDER-SAVE",
                    transdpwd_id: parseInt(idrecord),
                    transdpwd_tipedoc: transaksi_tipe_field,
                    transdpwd_idmember: transaksi_idmember_field,
                    transdpwd_bank_in: parseInt(transaksi_bankin_id_field),
                    transdpwd_bank_out: parseInt(transaksi_bankout_id_field),
                    transdpwd_amount: parseInt(transaksi_amount_field),
                    transdpwd_ipaddress: "",
                    transdpwd_note: transaksi_note_field,
                    transdpwd_status: transaksi_status_field,
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
                myModal_newentry.hide();
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
    async function handleUpdateTransaksi(e) {
        let flag = true
        let msg = ""
        if(transaksi_id_field == ""){
            flag = false
            msg += "The Document is required\n"
        }
        if(transaksi_idmember_field == ""){
            flag = false
            msg += "The Member is required\n"
        }
        if(e == "REJECTED"){
            if(transaksi_note_field == ""){
                flag = false
                msg += "The Note is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/transdpwdupdate", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"PROVIDER-SAVE",
                    transdpwd_id: transaksi_id_field,
                    transdpwd_idmember: transaksi_idmember_field,
                    transdpwd_note: transaksi_note_field,
                    transdpwd_status: e,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                
                msgloader = json.message;
                myModal_newentry.hide();
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
        temp_credit_member = 0;
        transaksi_temp_field = "";
        transaksi_id_field = "";
        transaksi_idmember_field = "";
        transaksi_bankin_id_field = "";
        transaksi_bankin_info_field = "";
        transaksi_bankout_id_field = "";
        transaksi_bankout_info_field = "";
        transaksi_amount_field = 0;
        transaksi_note_field = "";
        transaksi_status_field = "";
        transaksi_create_field = "";
        transaksi_update_field = "";
    }
    function callFunction(event){
        switch(event.detail){
            case "NEWDEPOSIT":
                NewDataTransaksi("New","DEPOSIT");
                break;
            case "NEWWITHDRAW":
                NewDataTransaksi("New","WITHDRAW");
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
    const handleKeyboard_number = (e) => {
        if (isNaN(parseInt(e.target.value))) {
            return e.target.value = 0;
        }else{
            return e.target.value = parseInt(e.target.value);
        }
	}
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
<div class="container-fluid" style="margin-top: 70px;">
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
                            placeholder="Search DOCUMENT"
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
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TIPE</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">&nbsp;</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DOCUMENT</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DATE</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MEMBER</th>
                                <th NOWRAP width="25%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">FROM</th>
                                <th NOWRAP width="25%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TO</th>
                                <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">AMOUNT</th>
                                <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">BEFORE</th>
                                <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">AFTER</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        {#if rec.home_status == "PROCESS"}
                                        <i on:click={() => {
                                                // e,tipe,id,idmember,bank_in,bank_out,bank_in_info,bank_out_info,amount,status,note,create,update
                                                NewDataTransaksi("Edit",rec.home_tipedoc,
                                                rec.home_id, rec.home_idmember, rec.home_nmmember,
                                                rec.home_bank_in, rec.home_bank_out,rec.home_bank_in_info,rec.home_bank_out_info,
                                                rec.home_amount, rec.home_status, rec.home_note,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                        {/if}
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {rec.home_status}
                                        </span>
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_tipedoc}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_tipeuserdoc_css}">
                                            {rec.home_tipeuserdoc}
                                        </span>
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_date}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmmember}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_bank_out_info}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_bank_in_info}</td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.home_amount_css}">{new Intl.NumberFormat().format(rec.home_amount)}</td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:#3600ff;font-weight:bold;">{new Intl.NumberFormat().format(rec.home_before)}</td>
                                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:#3600ff;font-weight:bold;">{new Intl.NumberFormat().format(rec.home_after)}</td>
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
	modal_title="{transaksi_tipe_field} / {sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Jenis Transaksi</label>
                    <input class="form-control"
                        type="text"
                        value="{transaksi_tipe_field}"
                        disabled
                        placeholder="Jenis Transaksi"/>
                </div>
                {#if transaksi_tipe_field == "DEPOSIT"}
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">From - Member</label>
                        <div class="input-group mb-3">
                            <Input bind:value={transaksi_bankout_info_field}
                                class="required"
                                type="text"
                                disabled
                                placeholder="Member"/>
                            <button on:click={() => {
                                ShowInOut("OUT");
                                }} class="btn btn-primary" type="button" id="button-addon2">...</button>
                        </div>
                    </div>
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">From</label>
                        <select
                            bind:value="{transaksi_bankout_id_field}" 
                            name="bankid" id="bankid" 
                            class="required form-control">
                            {#each member_listbank as rec}
                            <option value="{rec.memberbank_id}">{rec.memberbank_info}</option>
                            {/each}
                        </select>
                    </div>
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">To - Agen</label>
                        <div class="input-group mb-3">
                            <Input bind:value={transaksi_bankin_info_field}
                                class="required"
                                type="text"
                                disabled
                                placeholder="Agen"/>
                            <button on:click={() => {
                                ShowInOut("IN");
                                }} class="btn btn-primary" type="button" id="button-addon2">...</button>
                        </div>
                    </div>
                {:else}
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">From - Agen</label>
                        <div class="input-group mb-3">
                            <Input bind:value={transaksi_bankout_info_field}
                                class="required"
                                type="text"
                                disabled
                                placeholder="Agen"/>
                            <button on:click={() => {
                                ShowInOut("OUT");
                                }} class="btn btn-primary" type="button" id="button-addon2">...</button>
                        </div>
                    </div>
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">To - Member</label>
                        <div class="input-group mb-3">
                            <Input bind:value={transaksi_bankin_info_field}
                                class="required"
                                type="text"
                                disabled
                                placeholder="Member"/>
                            <button on:click={() => {
                                ShowInOut("IN");
                                }} class="btn btn-primary" type="button" id="button-addon2">...</button>
                        </div>
                    </div>
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">To</label>
                        <select
                            bind:value="{transaksi_bankin_id_field}" 
                            name="bankid" id="bankid" 
                            class="required form-control">
                            {#each member_listbank as rec}
                            <option value="{rec.memberbank_id}">{rec.memberbank_info}</option>
                            {/each}
                        </select>
                    </div>
                {/if}
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Amount</label>
                    <Input
                        on:keyup={handleKeyboard_number} 
                        bind:value={transaksi_amount_field}
                        class="required"
                        style="text-align: right;"
                        type="text"
                        placeholder="Amount"/>
                    <div id="passwordHelpBlock" class="form-text" style="text-align: right;font-size:11px;">
                        {new Intl.NumberFormat().format(transaksi_amount_field)}
                    </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Note</label>
                    <textarea
                        bind:value={transaksi_note_field} 
                        class="form-control" 
                        style="resize: none;"
                        placeholder="Note" id="floatingTextarea"></textarea>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {transaksi_create_field}<br />
                        Update : {transaksi_update_field}
                    </div>
                </div>
                {/if}
            </div>
        </div>
        
        
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSaveTransaksi();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>


<Modal
	modal_id="modalformdeposit"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="DOCUMENT : {transaksi_id_field}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Jenis Transaksi</label>
                    <input class="form-control"
                        type="text"
                        value="{transaksi_tipe_field}"
                        disabled
                        placeholder="Jenis Transaksi"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">From - Member</label>
                    <div class="alert alert-secondary" style="padding:10px;" role="alert">
                       {transaksi_nmmember_field}<br />
                       {transaksi_bankout_info_field}
                    </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">To - Agen</label>
                    <div class="alert alert-secondary" style="padding:10px;" role="alert">
                       {transaksi_bankout_info_field}
                    </div>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Amount</label>
                    <Input
                        on:keyup={handleKeyboard_number} 
                        bind:value={transaksi_amount_field}
                        class="required"
                        style="text-align: right;"
                        type="text"
                        placeholder="Amount"/>
                    <div id="passwordHelpBlock" class="form-text" style="text-align: right;font-size:11px;">
                        {new Intl.NumberFormat().format(transaksi_amount_field)}
                    </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Note</label>
                    <textarea
                        bind:value={transaksi_note_field} 
                        class="form-control" 
                        style="resize: none;"
                        placeholder="Note" id="floatingTextarea"></textarea>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {transaksi_create_field}<br />
                        Update : {transaksi_update_field}
                    </div>
                </div>
                {/if}
            </div>
        </div>
        
        
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleUpdateTransaksi("APPROVED");
            }} 
            button_function="SAVE"
            button_title="Approve"
            button_css="btn-warning"/>
        <Button on:click={() => {
                handleUpdateTransaksi("REJECTED");
            }} 
            button_function="SAVE"
            button_title="Rejected"
            button_css="btn-danger"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalformwithdraw"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="DOCUMENT : {transaksi_id_field}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Jenis Transaksi</label>
                    <input class="form-control"
                        type="text"
                        value="{transaksi_tipe_field}"
                        disabled
                        placeholder="Jenis Transaksi"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">From - Agen</label>
                    <div class="input-group mb-3">
                        <Input bind:value={transaksi_bankout_info_field}
                            class="required"
                            type="text"
                            disabled
                            placeholder="Agen"/>
                        <button on:click={() => {
                            ShowInOut("OUT");
                            }} class="btn btn-primary" type="button" id="button-addon2">...</button>
                    </div>
                </div>
                
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">To - Member</label>
                    <div class="alert alert-secondary" style="padding:10px;" role="alert">
                        {transaksi_nmmember_field}<br />
                        {transaksi_bankin_info_field}
                     </div>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Amount</label>
                    <Input
                        on:keyup={handleKeyboard_number} 
                        bind:value={transaksi_amount_field}
                        class="required"
                        style="text-align: right;"
                        type="text"
                        placeholder="Amount"/>
                    <div id="passwordHelpBlock" class="form-text" style="text-align: right;font-size:11px;">
                        {new Intl.NumberFormat().format(transaksi_amount_field)}
                    </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Note</label>
                    <textarea
                        bind:value={transaksi_note_field} 
                        class="form-control" 
                        style="resize: none;"
                        placeholder="Note" id="floatingTextarea"></textarea>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {transaksi_create_field}<br />
                        Update : {transaksi_update_field}
                    </div>
                </div>
                {/if}
            </div>
        </div>
        
        
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSaveTransaksi();
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
    modal_footer={false}
    modal_search={true}>
    <slot:template slot="search">
        <div class="col-lg-12" style="padding: 5px;">
            <input
                bind:value={member_search}
                type="text"
                class="form-control"
                placeholder="Search Member"
                aria-label="Search"/>
        </div>
    </slot:template>
    <slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">MEMBER</th>
                    <th width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CREDIT</th>
                </tr>
            </thead>
            <tbody>
            {#each filter_listmember as rec}
                <tr on:click={() => {
                        InsertInOut(rec.member_id, rec.member_name, rec.member_listbank, rec.member_credit);
                    }} style="cursor: pointer;">
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.member_no}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.member_name}</td>
                    <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:#3600ff;font-weight:bold;">{new Intl.NumberFormat().format(rec.member_credit)}</td>
                </tr>
            {/each}
            </tbody>
        </table>
    </slot:template>
</Modal>

<Modal
    modal_id="modallistagenbank"
    modal_size="modal-dialog-centered"
    modal_title="AGEN BANK"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_footer={false}
    modal_search={true}>
    <slot:template slot="search">
        <div class="col-lg-12" style="padding: 5px;">
            <input
                bind:value={agenbank_search}
                type="text"
                class="form-control"
                placeholder="Search List Bank Agen"
                aria-label="Search"/>
        </div>
    </slot:template>
    <slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">LIST BANK AGEN</th>
                </tr>
            </thead>
            <tbody>
            {#each filter_listagenbank as rec}
                <tr on:click={() => {
                        InsertInOut(rec.agenbank_id, rec.agenbank_info,"",0);
                    }} style="cursor: pointer;">
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.agenbank_no}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.agenbank_info}</td>
                </tr>
            {/each}
            </tbody>
        </table>
    </slot:template>
</Modal>