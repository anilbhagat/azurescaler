az group create --name myResourceGroup --location westus
az storage account create --name httptriggerazurescaler3 --resource-group myResourceGroup --sku Standard_LRS --location westus 

az functionapp create --consumption-plan-location westus --runtime custom --functions-version 4 --name HttpTriggerAzureScaler --os-type linux --resource-group myResourceGroup -s httptriggerazurescaler3

func azure functionapp publish HttpTriggerAzureScaler