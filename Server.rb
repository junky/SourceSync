require 'webrick'
include WEBrick

class SourceCodeServlet < WEBrick::HTTPServlet::AbstractServlet

  @@path_prefix = "./temp"
  
  def do_GET(request, response)
    response.status = 200
    response['Content-Type'] = "text/plain"
    response.body = "WeeebRick"
  end

  def parse_query(content_type, body)
    boundary = content_type.match(/^multipart\/form-data; boundary=(.+)/)[1]
    boundary = HTTPUtils::dequote(boundary)
    return HTTPUtils::parse_form_data(body, boundary)
  end

  def do_POST(request, response)
    body = String.new
    request.body do |chunk|
      body << chunk
    end

    filedata = parse_query(request['content-type'], body)
    path = filedata['path']
    if(not @@path_prefix.nil? and not @@path_prefix.empty?)
      path = @@path_prefix + path[1, path.size-1]
    end
    content_file = filedata['upload']
      
#    puts @@path_prefix
#    puts "----------"
#    puts content_file
#    puts "----------"
#    content_file.delete("\r").each_byte {|c| print c, ' ' }
    
    FileUtils.mkdir_p(File.dirname(path))
          
    file = File.new(path, "w")
    file.write(content_file.delete("\r"))
    file.close()
                
    response.status = 200
    response['Content-Type'] = "text/plain"
    response.body = "OK"
  end
end

server = HTTPServer.new(
      :Port            => 8000,
      :DocumentRoot    => './'
    )
server.mount "/", SourceCodeServlet

trap("INT"){ server.shutdown }

#WEBrick::Daemon.start
server.start

#https://github.com/betten/SoundCloud-Developer-Challenge/blob/master/server.rb
#
#