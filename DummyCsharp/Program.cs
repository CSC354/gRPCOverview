using System;
using System.Threading;
using System.Threading.Tasks;
using Csharp_Services;
using Grpc.Net.Client;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;


using var channel = GrpcChannel.ForAddress("http://localhost:4040");
var client = new RouteGuide.RouteGuideClient(channel);
var myhello = client.GetHello(new Hello() {Hello_ = "Csharp"});
Console.Write(myhello.Hello_);