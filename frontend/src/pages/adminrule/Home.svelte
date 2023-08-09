<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import Button from "../../components/Button.svelte";
    import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";

    export let table_header_font;
    export let table_body_font;
    export let token;
    export let listAdminrule = [];
    export let totalrecord = 0;
    let dispatch = createEventDispatcher();
    let title_page = "Admin Rule";
    let sData = "";
    let myModal_newentry = "";
    let css_loader = "display: none;";
    let msgloader = "";

    const schema = yup.object().shape({
        name: yup
            .string()
            .required("The Rule is required")
            .matches(/^[a-z0-9]+$/, "Rule must Character a-z or 1-9 ")
            .min(3, "The Rule minimal 3 Character")
            .max(30, "The Rule mmaximal 30 Character"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            name: "",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            handleSave(values.name);
        },
    });
    const NewData = () => {
        clearField();
        sData = "New";
        myModal_newentry = new bootstrap.Modal(
            document.getElementById("modalentry")
        );
        myModal_newentry.show();
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const EditData = (e, f) => {
        const adminrule = {
            e,
            f,
        };
        dispatch("handleEditData", adminrule);
    };
    async function handleSave(name) {
        const res = await fetch("/api/saveadminrule", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "ADMINRULE-SAVE",
                adminrule_idadmin: name,
                adminrule_rule: "",
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            msgloader = json.message;
            myModal_newentry.hide();
            RefreshHalaman();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    function clearField() {
        $form.username = "";
        $form.password = "";
        $form.rule = "";
        $form.name = "";
    }
    $: {
        if ($errors.name) {
            alert($errors.name);
            $form.name = "";
        }
    }

    function callFunction(event) {
        switch (event.detail) {
            case "NEW":
                NewData();
                break;
            case "REFRESH":
                RefreshHalaman();
                break;
            case "SAVE":
                handleSubmit();
                break;
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
                button_css="btn-dark"
            />
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"
            />
            <Panel card_title={title_page} card_footer={totalrecord}>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">RULE</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                            <tbody>
                                {#each listAdminrule as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i on:click={() => {
                                                    EditData(
                                                        rec.adminrule_idadmin,
                                                        rec.adminrule_rule
                                                    );
                                                }} class="bi bi-pencil"/>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.adminrule_no}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.adminrule_idadmin}</td>
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
    modal_title={title_page + "/" + sData}
    modal_footer_css="padding:5px;"
    modal_footer={true}
>
    <slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <input
                on:change={handleChange}
                bind:value={$form.name}
                invalid={$errors.name.length > 0}
                type="text"
                class="form-control required"
                placeholder="Name"
                aria-label="Name"
            />
        </div>
    </slot:template>
    <slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"
        />
    </slot:template>
</Modal>
