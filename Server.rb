require 'webrick'
include WEBrick

class Simple < WEBrick::HTTPServlet::AbstractServlet
  
  def do_GET(request, response)
    response.status = 200
    response['Content-Type'] = "text/plain"
    response.body = "WeeebRick"
  end
end

server = HTTPServer.new(
      :Port            => 8000,
      :DocumentRoot    => './'
    )
server.mount "/", Simple

trap("INT"){ server.shutdown }

server.start
