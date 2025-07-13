param location string

@description('Name of the Vault')
param vaultName string = 'vault-${uniqueString(resourceGroup().id)}'

@description('The SKU family')
@allowed(['A'])
param skuFamily string = 'A'

@description('The SKU name')
@allowed(['standard'])
param skuName string = 'standard'

resource keyVault 'Microsoft.KeyVault/vaults@2024-11-01' = {
  name: vaultName
  location: location
  properties: {
    sku: {
      family: skuFamily
      name: skuName
    }
    tenantId: subscription().tenantId
    accessPolicies: []
  }
}
