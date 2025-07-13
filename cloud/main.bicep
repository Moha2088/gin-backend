targetScope = 'resourceGroup'

@description('Resourcegroup location.')
param location string = resourceGroup().location

module vault 'modules/keyvault.bicep'={
  name: 'KeyVault'
  params: {
    location: location
  }
}
