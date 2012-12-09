require 'rest_client'

RestClient.post( 'http://localhost:8000',
  :upload => { :file => File.new('C:\Temp\key.txt', 'rb')}
)