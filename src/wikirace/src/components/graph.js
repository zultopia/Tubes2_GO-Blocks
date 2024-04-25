import React from 'react';
import Graph from "react-graph-vis";

function Graf({paths}) {
    var tempNodes = []
    var tempEdges = []
    paths.map((path, index) => {
        var len = path.length
        for (var i = 0; i < len; i++) {
            var conditional = true;
            tempNodes.every((node, index) => {
                if (node.id === path[i].Title) {
                    conditional = false;
                    return false
                }
                return true
            })
            
            console.log(conditional)
            if (conditional) {
                tempNodes.push({id: path[i].Title, label: path[i].Title, title: path[i].URL})
            }
            if (i < len-1) {
                conditional = true;
                tempEdges.every((edge) => {
                    if (edge.from === path[i].Title && edge.to === path[i+1].Title) {
                        conditional = false;
                        return false
                    }
                    return true
                })
                if (conditional) {
                    tempEdges.push({from: path[i].Title, to: path[i+1].Title})
                }
            }
        }
    })
    const graph = {
      nodes: tempNodes,
      edges: tempEdges
    };
  
    const options = {
        nodes: {
            shape: "dot",
            size: 16,
          },
          physics: {
            enabled: true,
        //     forceAtlas2Based: {
        //       gravitationalConstant: -26,
        //       centralGravity: 0.005,
        //       springLength: 230,
        //       springConstant: 0.18,
        //     },
        //     maxVelocity: 146,
        //     solver: "forceAtlas2Based",
        //     timestep: 0.35,
        //     stabilization: { iterations: 150 },
          },
          interaction: {
            navigationButtons: true,
            zoomView: true, 
            hover: true,
          },
    };
  
    // const events = {
    //   select: function(event) {
    //     var { nodes, edges } = event;
    //   }
    // };
    return (
      <Graph
        graph={graph}
        options={options}
      />
    );
  }

  export default Graf;